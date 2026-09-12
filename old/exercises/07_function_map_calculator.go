package main

import (
	"fmt"
	"strconv"
)

func calcAdd(i int, j int) int      { return i + j }
func calcSubtract(i int, j int) int { return i - j }
func calcDiv(i int, j int) int      { return i / j }
func calcMul(i int, j int) int      { return i * j }

var calcOpMap = map[string]func(int, int) int{
	"+": calcAdd,
	"-": calcSubtract,
	"*": calcMul,
	"/": calcDiv,
}

func Run07FunctionMapCalculator() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := calcOpMap[op]
		if !ok {
			fmt.Println("unsupported operator: ", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
