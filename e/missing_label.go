package main

import "fmt"

func main() {
loop:

	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisble by 3 but not 2")
		case 7:
			fmt.Println("exit the loop")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}
