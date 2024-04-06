package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func sumMutex() {
	arr := []int{2, 4, 6, 8, 10}
	var sum int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			mu.Lock()
			sum = sum + (v * v)
		}(v)
	}

	wg.Wait()
	fmt.Println(sum)
}

func sumAtomic() {
	arr := []int{2, 4, 6, 8, 10}
	var sum atomic.Int64
	wg := sync.WaitGroup{}
	for _, value := range arr {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			sum.Add(int64(value * value))
		}(value)
	}
	wg.Wait()
	fmt.Println(sum.Load())
}
