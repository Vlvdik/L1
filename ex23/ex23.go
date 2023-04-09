package ex23

import (
	"fmt"
	"log"
)

func Test() {
	someSlice := []int{1, 2, 3, 4, 5, 6, 7, 9, 1}

	if len(someSlice) < 1 {
		log.Fatal("[main] received empty slice")
	}

	var i uint
	fmt.Print("[main] enter index of the item to be deleted: ")
	_, err := fmt.Scanln(&i)
	if err != nil {
		log.Panicf("error when trying to scan index: %s", err.Error())
	}

	if i > uint(len(someSlice)-1) {
		log.Fatalf("[main] %d is too big for this slice\nmax index: %d\n", i, len(someSlice)-1)
	}

	someSlice = append(someSlice[:i], someSlice[i+1:]...)
	fmt.Printf("[main] done\nresult: %v\n", someSlice)
}
