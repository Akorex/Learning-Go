package main

import (
	"fmt"
	"os"
	"strconv"
)

func Run11CliArgsAndDecimal() {
	if len(os.Args) < 3 {
		fmt.Println("Need two parameters: amount and percent")
		return
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error parsing amount:", err)
		return
	}

	percent, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Error parsing percent:", err)
		return
	}

	// Calculated percentage using standard library float64
	total := amount * (percent / 100.0)
	fmt.Printf("Amount: %.2f, Percent: %.2f%%, Result: %.2f\n", amount, percent, total)
}
