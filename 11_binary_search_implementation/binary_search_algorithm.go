package main

import (
	"fmt"
	"math"
)

type User struct {
	Name string
	Age  int
}

// Node contain value and left right node.
type Node struct {
	Arr   []User
	Left  *Node
	Right *Node
}

// BST or binary search tree is algorithm to find a target value from a sorted array using divide and conquer.
// if target is found, returning the index of it.
func main() {
	arr := []User{
		{
			Name: "John",
		},
		{
			Name: "Yoga",
		},
		{
			Name: "Budi",
		},
		{
			Name: "Bill",
		},
		{
			Name: "Maxwell",
		},
	}
	arr = mergeSort(arr)
	target := "Maxwell"
	index, err := iterativeBinarySearch(arr, target)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("target %s index location is %d", target, index)
	}
}

// input : sorted array of integer and target
// output : index of target's location.
func iterativeBinarySearch(arr []User, target string) (int, error) {
	n := len(arr)
	start := 0
	end := n - 1
	for {
		mid := int((start + end) / 2)
		if arr[mid].Name == target {
			return mid, nil
		}
		if target < arr[mid].Name {
			end = mid - 1
		} else if target > arr[mid].Name {
			start = mid + 1
		} else {
			break
		}
	}
	return 0, fmt.Errorf("not found")
}

func mergeSort(arr []User) []User {
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
	var arr []User
	for {
		if i == len(node.Left.Arr) {
			arr = append(arr, node.Right.Arr[j:]...)
			break
		}

		if j == len(node.Right.Arr) {
			arr = append(arr, node.Left.Arr[i:]...)
			break
		}

		if node.Left.Arr[i].Name <= node.Right.Arr[j].Name {
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
