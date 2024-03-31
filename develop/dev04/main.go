package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	in := make(chan interface{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer close(in)
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				in <- getRndType()
				time.Sleep(2 * time.Second)
			}
		}
	}()

	run(in, &wg, 5) // 5 вывести в аргументы
	wg.Wait()
}

func run(in <-chan interface{}, wg *sync.WaitGroup, numWorker int) {
	for i := 1; i < numWorker-1; i++ {
		wg.Add(1)
		go worker(in, i, wg)
	}

}

func worker(in <-chan interface{}, numGourtine int, wg *sync.WaitGroup) {
	for t := range in {
		fmt.Println(numGourtine, ": ", t)
	}
	wg.Done()
}

func getRndType() interface{} {
	i := rand.IntN(6)
	switch i {
	case 0:
		return 1
	case 1:
		return "string"
	case 2:
		return 1.2
	case 3:
		return true
	case 4:
		return []int{1, 2, 3}
	case 5:
		return map[interface{}]interface{}{1: 1}
	case 6:
		return func() {
			fmt.Println("its func))")
		}
	default:
		return nil
	}
}
