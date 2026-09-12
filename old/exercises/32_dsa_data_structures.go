package main

import (
	"errors"
	"fmt"
)

type DsaInteger interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func dsaDivAndRemainder[T DsaInteger](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("Division by zero")
	}
	return num / denom, num % denom, nil
}

type DsaQueue[T comparable] struct {
	vals []T
}

func (q *DsaQueue[T]) Enqueue(val T) {
	q.vals = append(q.vals, val)
}

func (q *DsaQueue[T]) Dequeue() (T, bool) {
	if len(q.vals) == 0 {
		var zero T
		return zero, false
	}
	first := q.vals[0]
	q.vals = q.vals[1:]
	return first, true
}

func (q *DsaQueue[T]) IsEmpty() bool {
	return len(q.vals) == 0
}

func (q *DsaQueue[T]) Length() int {
	return len(q.vals)
}

func (q *DsaQueue[T]) Contains(val T) bool {
	for _, v := range q.vals {
		if v == val {
			return true
		}
	}
	return false
}

type DsaStack[T comparable] struct {
	vals []T
}

func (s *DsaStack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *DsaStack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *DsaStack[T]) Length() int {
	return len(s.vals)
}

func (s *DsaStack[T]) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *DsaStack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func DsaMap[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func DsaReduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func DsaFilter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Run32DsaDataStructures() {
	q := DsaQueue[string]{}
	q.Enqueue("hello")
	q.Enqueue("world")
	fmt.Println("Queue length:", q.Length())

	st := DsaStack[int]{}
	st.Push(42)
	fmt.Println("Stack top popped:", func() int { v, _ := st.Pop(); return v }())

	qRem, rRem, err := dsaDivAndRemainder(10, 3)
	fmt.Println("divAndRemainder:", qRem, rRem, err)
}
