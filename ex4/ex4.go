package ex4

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func work(resCh chan struct{}) {
	ping := time.Duration(rand.Intn(4)) * time.Second
	time.Sleep(ping)
	resCh <- struct{}{}
}

func worker(ctx context.Context, i int, ch chan interface{}) {
	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Imitation of work process
	resCh := make(chan struct{})
	go work(resCh)
	fmt.Printf("worker %d: started\n", i)

	for {
		select {
		case <-resCh:
			ch <- i
			return
		case <-timeout.Done():
			fmt.Printf("Worker %d: timeout\n", i)
			return
		}
	}
}

func do() {
	var n int
	fmt.Print("Enter the num of workers: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Panicf("error when trying to scan num of workers: %s\n", err.Error())
	}

	// Use a variable for the number of jobs.
	// For simplicity, let the number of jobs = number of workers.
	jobs := n

	// Randomize the generator
	rand.Seed(time.Now().UnixNano())

	// Parent context
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	// Using this channel, we simulate receiving important data.
	// In this case - the counter variable from the loop below
	ch := make(chan interface{})

	for i := 1; i <= n; i++ {
		go worker(ctx, i, ch)
	}

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

loop:
	for jobs != 0 {
		select {
		case <-gracefulShutdown:
			fmt.Println("main program: DONE BY CTRL + C")
			return
		case i := <-ch:
			fmt.Printf("worker %v: finished the job\n", i)
			jobs--
		case <-ctx.Done():
			fmt.Println("main program: timeout")
			break loop
		}

		if jobs == 0 {
			fmt.Println("main program: all of the workers have completed")
			break
		}
	}

	fmt.Println("main program: awaiting signal")
	<-gracefulShutdown
	fmt.Println("main program: DONE BY CTRL + C")
}

func Test() {
	do()
}
