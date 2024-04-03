package main

import "fmt"

func main() {
	fmt.Println(quicksort([]int{1, 22, 13, 0, 15, 302, -120}))
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	middle := arr[0]
	var left, right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] <= middle {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quicksort(left), middle), quicksort(right)...)
}
