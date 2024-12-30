package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func main() {
	n := 0
	start := time.Now()

	for i := 0; i < 100_000_000; i++ {
		n += 1
	}

	elapsed := time.Since(start)

	fmt.Println(elapsed)

	fmt.Println(addTo(3))

	x, y, z := divAndRemaninder(5, 2)

	fmt.Println(x, y, z)

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)

}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func divAndRemaninder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("Cannot divide by zero")
		return result, remainder, err
	}

	result, remainder = num/denom, num%denom

	return result, remainder, err
}
