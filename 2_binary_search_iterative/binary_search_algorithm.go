package main

import (
	"fmt"
	// "math"
)

// BST or binary search tree is algorithm to find a target value from a sorted array using divide and conquer.
// if target is found, returning the index of it.
func main() {

	arr := []int{1, 3, 7, 8, 11, 12, 15, 16, 21, 25}
	target := 21
	index, err := recursiveBinarySearch(arr, target, 0, len(arr)-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("target %d index location is %d", target, index)
	}
}

// input : sorted array of integer and target
// output : index of target's location.
func recursiveBinarySearch(arr []int, target, start, end int) (int, error) {
	mid := int((start + end) / 2)
	if arr[mid] == target {
		return mid, nil
	}
	if target < arr[mid] {
		return recursiveBinarySearch(arr, target, start, mid-1)
	} else if target > arr[mid] {
		return recursiveBinarySearch(arr, target, mid+1, end)
	}
	return 0, fmt.Errorf("not found")
}
