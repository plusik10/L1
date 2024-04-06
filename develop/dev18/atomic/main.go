package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	atomic.Int64
}

func (c *Counter) Inc() {
	c.Add(1)
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.Load())
}
