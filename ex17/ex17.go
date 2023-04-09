package ex17

import "fmt"

func binarySearch(value int, sortArr []int) (isHere bool, idx int) {
	l, r := 0, len(sortArr)

	if len(sortArr) < 1 {
		return false, -1
	}

	if value < sortArr[l] || value > sortArr[r-1] {
		return false, -1
	}

	for l < r {
		mid := (l + r) / 2

		if value == sortArr[mid] {
			return true, mid
		}

		if value < sortArr[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if sortArr[r] != value {
		return false, -1
	} else {
		return true, r
	}
}

func Test() {
	arr := []int{1, 3, 6, 9, 12}
	isHere, idx := binarySearch(6, arr)
	fmt.Printf("[main] {left side search} done\nsearching value: %d, isHere: %t, index: %d\n", 6, isHere, idx)
	isHere, idx = binarySearch(7, arr)
	fmt.Printf("[main] {left side search} done\nsearching value: %d, isHere: %t, index: %d\n", 7, isHere, idx)
}
