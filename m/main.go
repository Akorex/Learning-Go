package main

import (
	"fmt"
)

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Square(val int) int {
	return val * val
}

func main() {
	words := []string{"One", "Potato", "Two", "Potato"}
	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)
	lengths := Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)
	sum := Reduce(lengths, 0, func(acc int, val int) int {
		return acc + val
	})
	fmt.Println(sum)

	squared := Map([]int{1, 2, 3, 4, 5}, Square)

	result := Map([]int{1, 2, 3, 4}, func(v int) int { return v * v })

	fmt.Println(squared)
	fmt.Println(result)

}
