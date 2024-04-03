package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 99}
	fmt.Println(search(a, 1))

}

func search(arr []int, findElem int) int {
	l := 0
	h := len(arr) - 1
	for l <= h {
		middle := (l + h) / 2
		guess := arr[middle]
		if guess == findElem {
			return middle
		}
		if guess < findElem {
			l = middle + 1
		} else {
			h = middle - 1
		}
	}
	return -1
}
