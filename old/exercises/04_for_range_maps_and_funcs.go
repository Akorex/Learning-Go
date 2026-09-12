package main

import (
	"errors"
	"fmt"
)

func addToForRange(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func divAndRemainderForRange(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by 0")
	}

	return num / denom, num % denom, nil
}

func Run04ForRangeMapsAndFuncs() {
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilna": false}

	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}

	ansSlice := addToForRange(3)
	fmt.Println(ansSlice)

	ansSlice = addToForRange(3, 2)
	fmt.Println(ansSlice)

	ansSlice = addToForRange(3, 2, 4, 6, 8)
	fmt.Println(ansSlice)

	a := []int{4, 3}
	ansSlice = addToForRange(3, a...)
	fmt.Println(ansSlice)

	result, remainder, err := divAndRemainderForRange(5, 2)
	fmt.Println(result, remainder, err)

	count, err := fmt.Println("Hello, world")
	fmt.Println(count)
	fmt.Println(err)

	f1 := func(s string) int { return len(s) }
	f2 := func(s string) int {
		total := 0
		for _, v := range s {
			total += int(v)
		}
		return total
	}

	var myFuncVariable func(string) int
	myFuncVariable = f1
	ans := myFuncVariable("hello")
	fmt.Println(ans)

	myFuncVariable = f2
	ans = myFuncVariable("hello")
	fmt.Println(ans)
}
