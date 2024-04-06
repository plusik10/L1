package main

import (
	"fmt"
	"log"
)

func main() {
	slices := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s, err := del(slices, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}

func del(s []int, inx int) ([]int, error) {
	if len(s) < inx {
		return nil, fmt.Errorf("index out of range")
	}

	return append(s[:inx], s[inx+1:]...), nil
}
