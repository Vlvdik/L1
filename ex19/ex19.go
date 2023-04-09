package ex19

import "fmt"

func reverseString(s []rune) string {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	return string(s)
}

func Test() {
	s := "some string aboba"
	fmt.Printf("[main] received string: %s\n", s)
	reversed := reverseString([]rune(s))
	fmt.Printf("[main] done\nreversed string: %s\n", reversed)
}
