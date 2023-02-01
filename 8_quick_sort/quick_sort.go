package main

import "fmt"

func main() {
	arr := []int{3, 7, 2, 5, 1, 4, 0, 6}
	quickSort(&arr, 0, len(arr)-1)

	fmt.Println(arr)
}

// partition
func partition(arr *[]int, start, end int) int {

	pIndex := end

	for i := start; i < end; i++ {
		if (*arr)[i] >= (*arr)[pIndex] {
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[pIndex]
			(*arr)[pIndex] = temp
			pIndex = i
		}
	}
	return pIndex
}

func quickSort(arr *[]int, start, end int) {
	if start >= end || start < 0 {
		return
	}

	pivot := partition(arr, start, end)

	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)

}
