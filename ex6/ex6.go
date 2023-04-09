package ex6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func timer(quit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// imitation of user requirements
	time.Sleep(3 * time.Second)
	quit <- struct{}{}
}

func chanGoroutine(quit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-quit:
			fmt.Println("[goroutine (with channel)] recieved data from channel")
			return
		}
	}
}

func withChannel() {
	quit := make(chan struct{})
	defer close(quit)

	var wg sync.WaitGroup
	wg.Add(2)

	go chanGoroutine(quit, &wg)
	go timer(quit, &wg)

	wg.Wait()

	fmt.Println("[main (with channel)] goroutine was cancelled")
}

func isCloseGoroutine(ok chan bool) {
	ok <- true
}

func ctxGoroutine(ctx context.Context, ok chan bool) {
	defer isCloseGoroutine(ok)

	// imitation of working process
	time.Sleep(10 * time.Second)
	ctx.Done()
}

func withoutChannel() {
	ok := make(chan bool)
	defer close(ok)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go ctxGoroutine(ctx, ok)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[main (with context)] timeout\n[goroutine (with context)] closed = %t", <-ok)
			return
		}
	}
}

func Test() {
	withChannel()
	withoutChannel()
}
