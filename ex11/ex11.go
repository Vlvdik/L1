package ex11

import "fmt"

func Test() {
	a := []string{"cat", "dog", "рыба", "получеловек", "cat"}
	b := []string{"dog", "корова", "получеловек", "dog"}
	m := make(map[string]struct{})
	res := make(map[string]struct{})

	for _, v := range a {
		m[v] = struct{}{}
	}

	// ok1 = true will mean that the element is contained in both sets
	// ok2 = true - check for non-content in the result set
	for _, v := range b {
		_, ok1 := m[v]
		if ok1 {
			_, ok2 := res[v]
			if !ok2 {
				res[v] = struct{}{}
			}
		}
	}

	fmt.Printf("[main] done\nresult: %v\n", res)
}
