package main

import (
	"fmt"
	"sort"
)

type SortedPerson struct {
	FirstName string
	LastName  string
	Age       int
}

type MethodAdder struct {
	start int
}

func (p SortedPerson) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func (a MethodAdder) AddTo(val int) int {
	return a.start + val
}

func Run09StructMethodsAndSorting() {
	people := []SortedPerson{
		{"Pat", "Peterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	p := SortedPerson{
		FirstName: "Adewole",
		LastName:  "Akorede",
		Age:       23,
	}

	output := p.String()
	fmt.Println(output)
	fmt.Println(p)

	myAdder := MethodAdder{start: 6}
	result := myAdder.AddTo(6)
	fmt.Println(result)

	f1 := myAdder.AddTo
	fmt.Println(f1(10))
}
