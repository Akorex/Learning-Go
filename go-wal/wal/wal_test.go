package wal_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/Akorex/Learning-Go/go-wal/wal"
)

func TestRoundTrip(t *testing.T) {
	// Create a safe temporary directory that cleans up after the test
	tmpDir := t.TempDir()
	path := tmpDir + "/test.wal"

	w, err := wal.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	// 1. Append a few mock care-request updates
	entries := [][]byte{
		[]byte("assign:user:42"),
		[]byte("confirm:booking:17"),
		[]byte("cancel:booking:9"),
	}

	for _, e := range entries {
		_, err := w.Append(wal.RecordWrite, e)
		if err != nil {
			t.Fatal(err)
		}
	}
	w.Close()

	// 2. Re-open the WAL (simulating an app restart)
	w2, err := wal.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer w2.Close()

	// 3. Read everything back to ensure it's there
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	reader := wal.NewReader(f)
	for i, expected := range entries {
		lsn, rtype, payload, err := reader.Next()
		if err != nil {
			t.Fatalf("failed reading entry %d: %v", i, err)
		}

		if lsn != uint64(i+1) {
			t.Errorf("expected LSN %d, got %d", i+1, lsn)
		}
		if rtype != wal.RecordWrite {
			t.Errorf("expected RecordWrite type")
		}
		if !bytes.Equal(payload, expected) {
			t.Errorf("expected payload %q, got %q", expected, payload)
		}
	}
}

func TestCrashRecovery(t *testing.T) {
	tmpDir := t.TempDir()
	path := tmpDir + "/crash.wal"

	w, _ := wal.Open(path)
	w.Append(wal.RecordWrite, []byte("good-record-1"))
	w.Append(wal.RecordWrite, []byte("good-record-2"))
	w.Close()

	// Simulate a crash mid-write: open the file raw and append a corrupt fragment
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	// Write a bad header fragment (just a few raw, incomplete bytes)
	f.Write([]byte{0x01, 0x00, 0x00, 0x00, 0xFF, 0xFF})
	f.Close()

	// Re-open the WAL. It should look at the garbage at the end,
	// realize it's invalid, and stop calmly without crashing.
	w2, err := wal.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer w2.Close()

	// Open raw file to read back what the reader parses
	f2, _ := os.Open(path)
	defer f2.Close()

	reader := wal.NewReader(f2)

	// We should get our first two good records perfectly
	_, _, p1, err := reader.Next()
	if err != nil || string(p1) != "good-record-1" {
		t.Fatalf("failed to recover record 1")
	}

	_, _, p2, err := reader.Next()
	if err != nil || string(p2) != "good-record-2" {
		t.Fatalf("failed to recover record 2")
	}

	// The third attempt should hit our corrupt trailing bytes and return an error
	_, _, _, err = reader.Next()
	if err == nil {
		t.Fatal("expected an error when hitting the corrupted crash boundary, but got nil")
	}
}
