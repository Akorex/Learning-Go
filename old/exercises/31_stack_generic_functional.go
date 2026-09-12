package main

import "fmt"

type FunctionalStack[T comparable] struct {
	vals []T
}

func (s *FunctionalStack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *FunctionalStack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *FunctionalStack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func FunctionalFilter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func FunctionalMap[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func FunctionalReduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Run31StackGenericFunctional() {
	var intStack FunctionalStack[int]

	intStack.Push(1)
	intStack.Push(30)
	intStack.Push(50)

	v, ok := intStack.Pop()
	fmt.Println("Popped:", v, ok)
	fmt.Println("Contains 30:", intStack.Contains(30))
	fmt.Println("Contains 10:", intStack.Contains(10))

	nums := []int{1, 2, 3, 4, 5}
	evens := FunctionalFilter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println("Evens:", evens)
}
