package main

import "fmt"

func div60Recover(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered in div60Recover:", v)
		}
	}()

	fmt.Println(60 / i)
}

func Run23PanicRecoverDivide() {
	for _, val := range []int{1, 2, 0, 6} {
		div60Recover(val)
	}
}
