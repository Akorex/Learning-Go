package main

import "fmt"

type IntQueue struct {
	vals []int
}

func (q *IntQueue) Push(v int) {
	q.vals = append(q.vals, v)
}

func (q *IntQueue) Top() int {
	if len(q.vals) == 0 {
		return -1
	}
	return q.vals[0]
}

func (q *IntQueue) Dequeue() int {
	if len(q.vals) == 0 {
		return -1
	}
	ans := q.vals[0]
	q.vals = q.vals[1:]
	return ans
}

func Run30QueueSlice() {
	var q IntQueue
	q.Push(10)
	q.Push(20)
	fmt.Println("Top:", q.Top())
	fmt.Println("Dequeued:", q.Dequeue())
	fmt.Println("Next:", q.Dequeue())
}
