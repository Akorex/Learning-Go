package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s %d", p.FirstName, p.LastName, p.Age)
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	p := Person{FirstName: "Akorede", LastName: "Adewole", Age: 24}

	fmt.Println(p.String())

	c := Counter{total: 9, lastUpdated: time.Now()}

	fmt.Println(c)

	fmt.Println(c.String())

}
