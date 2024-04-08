package main

import "fmt"

func main() {
	arrStr := []string{"abcd", "abCdefAaF", "aabcd", "asdfqwerJIPM"}
	for _, str := range arrStr {
		fmt.Println(str, " - ", check(str))
	}
}

func check(str string) bool {
	m := make(map[rune]struct{})
	for _, s := range str {
		if _, ok := m[s]; ok {
			return false
		} else {
			m[s] = struct{}{}
		}
	}
	return true
}
