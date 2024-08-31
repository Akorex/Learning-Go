package main

import "fmt"

func Square(val int) int {
	return val * val
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func main() {
	squared := Map([]int{1, 2, 3, 4, 5}, Square)
	fmt.Println(squared) // Should output: [1 4 9 16 25]
}
