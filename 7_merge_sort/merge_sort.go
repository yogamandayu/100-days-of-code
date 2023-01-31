package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{0, 5, 4, 7, 3, 2, 2, 8, 1}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
}

// Node contain value and left right node.
type Node struct {
	Arr   []int
	Left  *Node
	Right *Node
}

func mergeSort(arr []int) []int {
	node := &Node{
		Arr: arr,
	}
	divide(node)
	conquer(node)

	return node.Arr
}

// divide is to divide node until array only have 1 value or until the node leaf.
func divide(node *Node) {
	if len(node.Arr) == 1 {
		return
	}

	mid := int(math.Ceil(float64(len(node.Arr)) / 2))
	node.Left = &Node{
		Arr: node.Arr[0:mid],
	}
	divide(node.Left)

	node.Right = &Node{
		Arr: node.Arr[mid:],
	}
	divide(node.Right)

}

// conquer is to compare and merge left and right node.
func conquer(node *Node) {
	if node.Left != nil {
		conquer(node.Left)
		conquer(node.Right)
	} else {
		return
	}

	var i, j int
	var arr []int
	for {
		if i == len(node.Left.Arr) {
			arr = append(arr, node.Right.Arr[j:]...)
			break
		}

		if j == len(node.Right.Arr) {
			arr = append(arr, node.Left.Arr[i:]...)
			break
		}

		if node.Left.Arr[i] <= node.Right.Arr[j] {
			arr = append(arr, node.Left.Arr[i])
			i++
			continue
		}
		arr = append(arr, node.Right.Arr[j])
		j++
	}

	node.Arr = arr
	return
}
