package dev16

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	middle := arr[len(arr)/2]
	var left, right []int

	for i := 0; i < len(arr); i++ {
		if arr[i] <= middle {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quicksort(left), middle), quicksort(right)...)
}
