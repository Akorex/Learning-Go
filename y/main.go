package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

func main() {
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

	sort.Slice(people, func(i, j int) bool{
		return people[i].FirstName < people[j].FirstName
	})

	fmt.Println(people)


	slices.SortFunc(people, func(a, b Person) int{
		return a.Age - b.Age
	})

	fmt.Println(people)

	// sort by name
	slices.SortFunc(people, func(a, b Person) int{
		return cmp.Compare(a.FirstName, b.FirstName)
	})

	fmt.Println(people)
}