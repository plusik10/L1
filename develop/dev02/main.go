package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, num := range numbers {
		wg.Add(1)
		go func() {
			// go 1.22 выведет все как нужно, в 1.21 и ниже
			// придеться пробрасывать в функцию переменную
			fmt.Println(num * num)
			wg.Done()
		}()
	}
	wg.Wait()
}
