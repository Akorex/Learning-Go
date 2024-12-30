package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbDate = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	start := time.Now()
	for i := 0; i < len(dbDate); i++ {
		wg.Add(1)
		go dbCall(i)
	}

	wg.Wait()
	fmt.Printf("\nTotal Execution time: %v", time.Since(start))
	fmt.Printf("\nThe results are: %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000

	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is :", dbDate[i])
	m.Lock()
	results = append(results, dbDate[i])
	m.Unlock()
	wg.Done()
}
