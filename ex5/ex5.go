package ex5

import (
	"context"
	"fmt"
	"log"
	"time"
)

func work(ch chan interface{}) {
	fmt.Printf("[work] recieve data: %v\n", <-ch)
}

func do() {
	ch := make(chan interface{})
	defer close(ch)

	for i := 0; ; i++ {
		// The delay is added to reduce the number of messages received
		// and does not carry any useful load
		time.Sleep(time.Millisecond)

		go work(ch)
		ch <- i
	}
}

func Test() {
	var n uint
	fmt.Print("Enter the number of seconds for completion: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Panicf("error when trying to scan num of seconds: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	go do()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[main] timeout")
			return
		}
	}
}
