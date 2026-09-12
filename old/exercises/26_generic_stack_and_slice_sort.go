package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type GenericStack[T comparable] struct {
	vals []T
}

func (s *GenericStack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *GenericStack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *GenericStack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func (s *GenericStack[T]) Length() int {
	return len(s.vals)
}

func StackMap[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Run26GenericStackAndSliceSort() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{FirstName: "Akorede", LastName: "Adewole", Age: 25},
		{FirstName: "Adedeji", LastName: "Agbaoye", Age: 27},
		{FirstName: "Omotolani", LastName: "Adejumo", Age: 25},
	}

	fmt.Println(people)

	// sort by first name
	sort.Slice(people, func(i, j int) bool {
		return people[i].FirstName < people[j].FirstName
	})
	fmt.Println(people)

	slices.SortFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	fmt.Println(people)

	// sort by name
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.FirstName, b.FirstName)
	})
	fmt.Println(people)

	s := GenericStack[int]{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Stack length:", s.Length())
}
