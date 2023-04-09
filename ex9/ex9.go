package ex9

import "fmt"

func getSquare(in, out chan int) {
	num := <-in
	out <- num * num
}

func Test() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	in := make(chan int, 10)
	out := make(chan int, 10)

	for _, v := range a {
		in <- v
	}

	for range a {
		go getSquare(in, out)
	}

	for i := 1; i <= 10; {
		select {
		case res := <-out:
			fmt.Printf("[main] goroutine %d done\nresult: %d\n", i, res)
			i++
		}
	}

	close(in)
	close(out)
	fmt.Println("[main] done")
}
