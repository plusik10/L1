package main

import (
	"strings"
)

func someFunc() string {
	return createHugeString(1 << 10)[:1024]
}

func createHugeString(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteRune('o')
	}
	return sb.String()
}

func main() {
	someFunc()
}
