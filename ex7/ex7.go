package ex7

import (
	"fmt"
	"sync"
)

var m = make(map[int]struct{})

func write(i int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	m[i] = struct{}{}
	fmt.Printf("[goroutine %d] done\n", i)
	mu.Unlock()
}

func Test() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go write(i, &mu, &wg)
	}
	wg.Wait()

	fmt.Printf("[main] all goroutines are done\nresult: %v", m)
}
