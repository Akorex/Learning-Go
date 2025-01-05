package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need two parameters: amount and percent")
		os.Exit(1)
	}

	amount, err := decimal.NewFromString(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

}
