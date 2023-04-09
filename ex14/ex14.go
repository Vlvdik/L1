package ex14

import "fmt"

func Test() {
	a := 1
	b := "some string"
	c := true
	d := make(chan int)

	data := []interface{}{a, b, c, d}

	// We can expand the number of supported types by the same analogy
	for i, v := range data {
		switch v.(type) {
		case int:
			fmt.Printf("[main] value index = %d: its a integer! value: %v\n", i, v)
		case string:
			fmt.Printf("[main] value index = %d: its a string! value: %v\n", i, v)
		case bool:
			fmt.Printf("[main] value index = %d: its a bool! value: %v\n", i, v)
		case chan int:
			fmt.Printf("[main] value index = %d: its a integer channel! value: %v\n", i, v)
		}
	}
}
