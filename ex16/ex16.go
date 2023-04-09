package ex16

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	bigger, less := make([]int, 0), make([]int, 0)
	pivot := arr[0]

	for _, v := range arr[1:] {
		if v > pivot {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}

	arr = append(quickSort(less), pivot)
	arr = append(arr, quickSort(bigger)...)

	return arr
}

func partition(arr []int, l, r int) ([]int, int) {
	pivot := arr[r]
	i := l
	for j := l; j < r; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return arr, i
}

func classic(arr []int, l, r int) []int {
	if l < r {
		var pivot int
		arr, pivot = partition(arr, l, r)
		arr = classic(arr, l, pivot-1)
		arr = classic(arr, pivot+1, r)
	}
	return arr
}

func Test() {
	arr1 := []int{23, 1, 43, 0, 0, -3, 12, 11, 12, 12}
	fmt.Printf("[main] {with allocation of additional memory} \narray before sort: %v\n", arr1)
	fmt.Printf("[main] {with allocation of additional memory} quicksort was completed\narray after sort: %v\n", quickSort(arr1))

	arr2 := []int{23, 1, 43, 0, 0, -3, 12, 11, 12, 12}
	fmt.Printf("[main] {without allocation of additional memory} \narray before sort: %v\n", arr2)
	fmt.Printf("[main] {without allocation of additional memory} quicksort was completed\narray after sort: %v\n", classic(arr2, 0, len(arr2)-1))
}
