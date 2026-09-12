package main

import "fmt"

func GenericMap[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func GenericReduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func GenericFilter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func GenericSquare(val int) int {
	return val * val
}

func Run15GenericsMapReduceFilter() {
	words := []string{"One", "Potato", "Two", "Potato"}
	filtered := GenericFilter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)

	lengths := GenericMap(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)

	sum := GenericReduce(lengths, 0, func(acc int, val int) int {
		return acc + val
	})
	fmt.Println(sum)

	squared := GenericMap([]int{1, 2, 3, 4, 5}, GenericSquare)
	result := GenericMap([]int{1, 2, 3, 4}, func(v int) int { return v * v })

	fmt.Println(squared)
	fmt.Println(result)
}
