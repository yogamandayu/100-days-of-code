package main

import "fmt"

func main() {
	arr := []string{"A", "B", "C", "D"}
	// if arr throw as args to print, print will accept as one object array as [[A B C D]].
	print(arr)
	// to handle this, just place '...' at behind of args like print(arr...). But of course it's not gonna work because print only accept interface{}
	// print(arr...) not work! because print only accept []interface not []string.

	// so that arraymust be parsing into []interface.

	newArr := parseIntoSliceOfInterface(arr)
	print(newArr...)
}

func print(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func parseIntoSliceOfInterface(args []string) []interface{} {
	var s []interface{}

	for _, v := range args {
		s = append(s, v)
	}

	return s
}
