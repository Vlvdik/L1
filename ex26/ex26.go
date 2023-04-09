package ex26

import (
	"fmt"
	"unicode"
)

func isUnicString(s string) bool {
	sym := make(map[rune]struct{})
	for _, r := range []rune(s) {
		l := unicode.ToLower(r)
		_, ok := sym[l]
		if ok {
			return false
		}

		sym[l] = struct{}{}
	}

	return true
}

func Test() {
	s1 := "abcd 123"                  // expected: true
	s2 := "СтавЬ лайк если любишь WB" // expected: false
	s3 := "aboba"                     // expected: false
	s4 := ""                          // expected: true
	fmt.Printf("[main] s1 ===> isUnicString: %t\n", isUnicString(s1))
	fmt.Printf("[main] s2 ===> isUnicString: %t\n", isUnicString(s2))
	fmt.Printf("[main] s3 ===> isUnicString: %t\n", isUnicString(s3))
	fmt.Printf("[main] s4 ===> isUnicString: %t\n", isUnicString(s4))
}
