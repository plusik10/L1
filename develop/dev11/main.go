package main

import "fmt"

func main() {

	set1 := map[string]struct{}{"cat": {}, "dog": {}, "fish": {}}
	set2 := map[string]struct{}{"man": {}, "dog": {}, "fish": {}}

	fmt.Println(intersection(set1, set2))
}

func intersection(set1 map[string]struct{}, set2 map[string]struct{}) map[string]struct{} {
	result := make(map[string]struct{})
	for key := range set1 {
		if _, ok := set2[key]; ok {
			result[key] = struct{}{}
		}
	}
	return result
}
