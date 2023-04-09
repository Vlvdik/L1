package ex2

import (
	"fmt"
	"sync"
)

func getSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print(num*num, " ")
}

func sequential() {
	var wg sync.WaitGroup
	nums := [5]int{2, 4, 6, 8, 10}

	for _, num := range nums {
		wg.Add(1)
		go getSquare(num, &wg)
		wg.Wait()
	}
}

func nonSequential() {
	var wg sync.WaitGroup
	arr := [5]int{2, 4, 6, 8, 10}

	for _, v := range arr {
		wg.Add(1)
		go getSquare(v, &wg)
	}
	wg.Wait()
}

func Test() {
	sequential()
	nonSequential()
}
