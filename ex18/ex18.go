package ex18

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	value int
}

func newCounter() *counter {
	return &counter{value: 0}
}

func (c *counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *counter) GetCounterValue() int {
	return c.value
}

func doIncrement(num int, c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Inc()
	fmt.Printf("[goroutine %d] done\n", num)
}

func Test() {
	var wg sync.WaitGroup
	c := newCounter()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doIncrement(i, c, &wg)
	}
	wg.Wait()

	fmt.Printf("[main] all goroutines are done\ncounter value: %d\n", c.GetCounterValue())
}
