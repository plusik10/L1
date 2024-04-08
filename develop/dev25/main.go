package main

import (
	"context"
	"log"
	"time"
)

func main() {
	sleepV1(2 * time.Second)
	log.Println("sleepV1 out")
	sleepV2(3 * time.Second)
	log.Println("sleepV2 out")
}

func sleepV1(t time.Duration) {
	<-time.After(t)
}

func sleepV2(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	<-ctx.Done()
}
