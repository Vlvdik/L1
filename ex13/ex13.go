package ex13

import "fmt"

func Test() {
	a, b := 1, 2
	fmt.Printf("[main] {before} a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("[main] {after} a: %d, b: %d\n", a, b)
}
