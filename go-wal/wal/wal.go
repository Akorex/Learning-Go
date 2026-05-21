package wal

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

const headerSize = 13

type RecordType byte
type Reader struct {
	r io.Reader
}

const (
	RecordWrite      RecordType = 0x01
	RecordDelete     RecordType = 0x02
	RecordCheckpoint RecordType = 0x03
)

type WAL struct {
	mu      sync.Mutex
	file    *os.File
	writer  *bufio.Writer
	nextLSN atomic.Uint64
}

func Open(path string) (*WAL, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		return nil, err
	}

	w := &WAL{
		file:   f,
		writer: bufio.NewWriterSize(f, 64*1024),
	}

	lsn, err := w.recoverMaxLSN()
	if err != nil {
		f.Close()
		return nil, err
	}
	w.nextLSN.Store(lsn + 1)

	return w, nil
}

func (w *WAL) Append(rtype RecordType, payload []byte) (uint64, error) {
	w.mu.Lock() // lock the mutex so no other goroutine can write at the same time
	defer w.mu.Unlock()

	lsn := w.nextLSN.Add(1) - 1

	// step 1 - format the header
	hdr := make([]byte, headerSize)
	binary.LittleEndian.PutUint64(hdr[0:8], lsn)
	hdr[8] = byte(rtype)
	binary.LittleEndian.PutUint32(hdr[9:13], uint32(len(payload)))

	// step 2 - compute the checksum
	checksum := crc32.ChecksumIEEE(append(hdr, payload...))
	csumBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(csumBuf, checksum)

	// step 3 - write to Go's memory buffer
	for _, chunk := range [][]byte{hdr, payload, csumBuf} {
		if _, err := w.writer.Write(chunk); err != nil {
			return 0, err
		}
	}

	// step 4: flush data from Go's memory to OS kernel's memory
	if err := w.writer.Flush(); err != nil {
		return 0, err
	}

	if err := w.file.Sync(); err != nil {
		return 0, err
	}

	return lsn, nil

}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

func (r *Reader) Next() (lsn uint64, rtype RecordType, payload []byte, err error) {
	hdr := make([]byte, headerSize)

	if _, err = io.ReadFull(r.r, hdr); err != nil {
		return 0, 0, nil, err
	}

	lsn = binary.LittleEndian.Uint64(hdr[0:8])
	rtype = RecordType(hdr[8])
	length := binary.LittleEndian.Uint32(hdr[9:13])

	payload = make([]byte, length)
	if _, err = io.ReadFull(r.r, payload); err != nil {
		return 0, 0, nil, err // truncated payload - crash happened mid-write
	}

	var storedCRC [4]byte
	if _, err = io.ReadFull(r.r, storedCRC[:]); err != nil {
		return 0, 0, nil, err // file ended before the checksum
	}

	// recompute the checksum
	computed := crc32.ChecksumIEEE(append(hdr, payload...))
	if computed != binary.LittleEndian.Uint32(storedCRC[:]) {
		err = fmt.Errorf("wal: CRC mismatch at LSN %d- partial write at crash boundary", lsn)
	}

	return lsn, rtype, payload, err

}

// recoverMaxLSN scans the log to find the highest fully-written LSN,
// stopping at the first checksum failure (the crash boundary).
func (w *WAL) recoverMaxLSN() (uint64, error) {
	// Start reading from the very beginning of the file
	if _, err := w.file.Seek(0, io.SeekStart); err != nil {
		return 0, err
	}

	// Wrap our file in a buffered reader for speed
	r := NewReader(bufio.NewReaderSize(w.file, 64*1024))
	var maxLSN uint64

	for {
		lsn, _, _, err := r.Next()

		if err == io.EOF {
			break // Clean end of file. We are done!
		}

		if err != nil {
			// This is the core of crash recovery. We hit a corrupted or partial write.
			// We DO NOT crash the program. We just stop reading and accept the records
			// we've successfully verified up to this point.
			break
		}

		maxLSN = lsn
	}

	// Critical: move the file pointer back to the very end of the file,
	// so the next Append() writes safely over any corrupted tail data.
	_, err := w.file.Seek(0, io.SeekEnd)
	return maxLSN, err
}

func (w *WAL) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// Ensure any final bytes sitting in the Go memory buffer hit the OS kernel
	if err := w.writer.Flush(); err != nil {
		return err
	}
	// Force the OS to sync everything to disk before we shut down
	return w.file.Close()
}
