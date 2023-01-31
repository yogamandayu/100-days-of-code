package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	arr := []int{0, 5, 4, 7, 3, 2, 2, 8, 1}
	sortedArr := mergeSort(arr)
	log.Println(sortedArr)
}

// Node contain value and left right node.
type Node struct {
	Top   *Node
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
	fmt.Println(node.Arr)
	if len(node.Arr) == 1 {
		return
	}

	mid := int(math.Ceil(float64(len(node.Arr)) / 2))
	fmt.Println(mid)
	node.Left = &Node{
		Top:   node,
		Arr:   node.Arr[0:mid],
		Left:  &Node{},
		Right: &Node{},
	}
	divide(node.Left)

	node.Right = &Node{
		Top:   node,
		Arr:   node.Arr[mid:],
		Left:  &Node{},
		Right: &Node{},
	}
	divide(node.Right)

}

// conquer is to compare and merge left and right node.
func conquer(node *Node) {
	if node.Left != nil {
		fmt.Println(node.Left.Arr)
		conquer(node.Left)
	}
	if node.Right != nil {
		fmt.Println(node.Right.Arr)
		conquer(node.Right)
	}
	if node.Top == nil {
		return
	}

	var arr []int
	var indexLeft, indexRight int
	for {
		fmt.Println("arr", arr)
		left := node
		right := node.Top.Right

		if indexLeft == len(left.Arr)-1 {
			arr = append(arr, right.Arr...)
			break
		}
		if indexRight == len(right.Arr)-1 {
			arr = append(arr, left.Arr...)
			break
		}
		if left.Arr[indexLeft] >= right.Arr[indexRight] {
			arr = append(arr, left.Arr[indexLeft])
			indexLeft++
			continue
		}
		if right.Arr[indexRight] >= left.Arr[indexLeft] {
			arr = append(arr, right.Arr[indexRight])
			indexRight++
			continue
		}
	}
	if node.Top == nil {
		return
	}
	log.Println(arr)
	node.Top.Arr = arr
	return
}
