package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	s := []int{1, 2, 3, 4, 5}

	target := 4

	fmt.Println(BinarySearch(s, target))
	fmt.Println(BinarySearch(s, 20))
}
