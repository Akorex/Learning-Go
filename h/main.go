package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Adder struct {
	start int
}

func (p Person) String() string {
	// (p Person) is the receiver specification
	// String() is the name of the method

	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {

	people := []Person{
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

	p := Person{
		FirstName: "Adewole",
		LastName:  "Akorede",
		Age:       23,
	}

	output := p.String()

	fmt.Println(output)

	myAdder := Adder{start: 6}

	result := myAdder.AddTo(6)
	fmt.Println(result)

	f1 := myAdder.AddTo

	fmt.Println(f1(10))

}
