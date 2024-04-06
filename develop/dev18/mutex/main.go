package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++

}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(c.value)
}
