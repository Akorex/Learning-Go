package main

import "fmt"

func div60Panic(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered from panic:", v)
		}
	}()

	fmt.Println(60 / i)
}

func Run19PanicDeferRecover() {
	for _, val := range []int{1, 2, 0, 6} {
		div60Panic(val)
	}
}
