package ex20

import (
	"fmt"
	"strings"
)

func reverseWords(wrds []string) string {
	l, r := 0, len(wrds)-1

	for l < r {
		wrds[l], wrds[r] = wrds[r], wrds[l]
		l++
		r--
	}

	return strings.Join(wrds, " ")
}

func Test() {
	s := "some string aboba"
	fmt.Printf("[main] received string: %s\n", s)
	reversed := reverseWords(strings.Split(s, " "))
	fmt.Printf("[main] done\nresult: %s\n", reversed)
}
