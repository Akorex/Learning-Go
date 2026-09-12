package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func timingAddTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func timingDivAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("Cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func Run25BenchmarkTimingAndSorting() {
	n := 0
	start := time.Now()
	for i := 0; i < 100_000_000; i++ {
		n += 1
	}
	elapsed := time.Since(start)
	fmt.Println("Loop elapsed:", elapsed)

	fmt.Println("Add to:", timingAddTo(3))

	x, y, z := timingDivAndRemainder(5, 2)
	fmt.Println("Division:", x, y, z)

	type BenchmarkPerson struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []BenchmarkPerson{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by age:", people)
}
