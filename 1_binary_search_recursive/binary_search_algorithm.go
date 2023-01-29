package main

import (
	"fmt"
	// "math"
)

// BST or binary search tree is algorithm to find a target value from a sorted array using divide and conquer.
// if target is found, returning the index of it.
func main() {

	arr := []int{1, 3, 7, 8, 11, 12, 15, 16, 21, 25}
	target := 11
	index, err := iterativeBinarySearch(arr, target)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("target %d index location is %d", target, index)
	}
}

// input : sorted array of integer and target
// output : index of target's location.
func iterativeBinarySearch(arr []int, target int) (int, error) {
	n := len(arr)
	start := 0
	end := n - 1
	for {
		mid := int((start + end) / 2)
		if arr[mid] == target {
			return mid, nil
		}
		if target < arr[mid] {
			end = mid - 1
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			break
		}
	}
	return 0, fmt.Errorf("not found")
}
