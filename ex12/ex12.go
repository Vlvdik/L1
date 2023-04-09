package ex12

import "fmt"

func Test() {
	in := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make(map[string]struct{})

	for _, v := range in {
		res[v] = struct{}{}
	}

	fmt.Printf("[main] done\nresult: %v", res)
}
