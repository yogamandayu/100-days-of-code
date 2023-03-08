package main

import "fmt"

/*
This code is for demonstrating user acceptance chaining.
User connect each other as a graph or linked list.
Acceptance will start from first node to last node.
Each turn will set IsAccepted into true.
The Objective is to set all node true.
You can't set a node to true if it's not that user turn.
To make it more complex, user's sequence have a type SERIAL, PARALLEL, or GROUP of SERIAL/PARALLEL.
*/

type User struct {
	Name       string
	IsAccepted bool
}

type Graph struct {
	Sequence string
}

func main() {
	fmt.Println("vim-go")
}
