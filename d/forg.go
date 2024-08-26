package main

import (
	"errors"
	"fmt"
	"os"
)

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func main() {
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilna": false}

	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	result, remainder, err := divAndRemaninder(5, 2)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result, remainder)

	// fmt.Println returns two values

	count, err := fmt.Println("Hello, world")

	fmt.Println(count)
	fmt.Println(err)

	var myFuncVariable func(string) int

	myFuncVariable = f1
	ans := myFuncVariable("hello")
	fmt.Println(ans)

	myFuncVariable = f2
	ans = myFuncVariable("hello")
	fmt.Println(ans)

}

func divAndRemaninder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by 0")
	}

	return num / denom, num % denom, nil
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0

	for _, v := range a {
		total += int(v)
	}

	return total
}
