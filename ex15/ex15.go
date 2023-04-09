package ex15

import (
	"log"
	"os"
)

// Source: https://ru.stackoverflow.com/questions/1446734/

func createHugeString(size int, f *os.File) string {
	return "very huge string (нет)"
}

func someFunc() []byte {
	f, err := os.Create("huge_string.txt")
	if err != nil {
		log.Panicf("error when trying to create huge_string.txt: %s", err.Error())
	}
	defer f.Close()

	// Passing a file to write a string into it
	return []byte(createHugeString(1<<10, f))[:100]
}

func Test() {
	someFunc()
}
