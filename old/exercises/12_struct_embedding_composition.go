package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s) ", e.Name, e.ID)
}

func Run12StructEmbeddingComposition() {
	m := Manager{
		Employee: Employee{
			Name: "Akorede",
			ID:   "1234",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID)
	fmt.Println(m.Description())
}
