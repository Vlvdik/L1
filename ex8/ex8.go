package ex8

import (
	"fmt"
	"log"
)

func Test() {
	var (
		mask int64 = 1
		num  int64
	)

	fmt.Println("Enter the num: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Panicf("error when trying to scan num: %s", err.Error())
	}

	var i uint8
	fmt.Println("Enter the shift")
	_, err = fmt.Scanln(&i)
	if err != nil {
		log.Panicf("error when trying to scan shift: %s", err.Error())
	}

	mask = mask << i

	fmt.Printf("DONE\nresult: %d", num^mask)
}
