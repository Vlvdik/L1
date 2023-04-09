package ex25

import (
	"fmt"
	"time"
)

func sleepAfter(duration time.Duration) {
	<-time.After(duration)
}

func sleepTicker(duration time.Duration) {
	ticker := time.Tick(duration)
	for range ticker {
		return
	}
}

func Test() {
	fmt.Println("[main] {with time.After} wait 3 seconds")
	sleepAfter(3 * time.Second)
	fmt.Println("[main] {with time.After} done")

	fmt.Println("[main] {with ticker} wait 3 seconds")
	sleepTicker(3 * time.Second)
	fmt.Println("[main] {with ticker} done")

}
