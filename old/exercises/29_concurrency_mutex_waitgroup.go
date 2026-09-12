package main

import (
	"fmt"
	"sync"
	"time"
)

func Run29ConcurrencyMutexWaitgroup() {
	var m sync.Mutex
	var wg sync.WaitGroup
	dbData := []string{"id1", "id2", "id3", "id4", "id5"}
	var results []string

	dbCall := func(i int) {
		var delay float32 = 100 // reduced delay for quick demo
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println("Result from database:", dbData[i])
		m.Lock()
		results = append(results, dbData[i])
		m.Unlock()
		wg.Done()
	}

	start := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}

	wg.Wait()
	fmt.Printf("Total Execution time: %v\n", time.Since(start))
	fmt.Printf("Results: %v\n", results)
}
