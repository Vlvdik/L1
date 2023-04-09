package ex3

import (
	"fmt"
	"sync"
)

func getSquare(num int, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*sum += num * num
}

func getSumOfSquare(nums ...int) int {
	var wg sync.WaitGroup
	var res int

	for _, num := range nums {
		wg.Add(1)
		go getSquare(num, &res, &wg)
	}
	wg.Wait()

	return res
}

func Test() {
	fmt.Printf("sum: %d\n", getSumOfSquare(2, 4, 6, 8, 10))
}
