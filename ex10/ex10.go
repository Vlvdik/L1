package ex10

import "fmt"

func Test() {
	res := make(map[int][]float32)
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0}

	for _, v := range temp {
		res[int(v)/10*10] = append(res[int(v)/10*10], v)
	}
	fmt.Printf("[main] done\nresult: %v\n", res)
}
