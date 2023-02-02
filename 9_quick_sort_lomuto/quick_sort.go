package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 8, 7, 0, 6, 1, 2, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]

	pIndex := lo

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}
	arr[hi], arr[pIndex] = arr[pIndex], arr[hi]
	return pIndex
}

func quickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(arr, lo, hi)
	fmt.Println(arr)

	quickSort(arr, lo, pivot-1)
	quickSort(arr, pivot+1, hi)
}
