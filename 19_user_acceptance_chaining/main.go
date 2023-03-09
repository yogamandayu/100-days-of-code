package main

import "fmt"

/*
This code is for demonstrating user acceptance chaining.
User connect each other as a graph or linked list.
Acceptance will start from first node to last node.
Each turn will set IsAccepted into true.
The Objective is to set all node true (depend on sequence).
You can't set a node to true if it's not user turn.
To make it more complex, user's sequence have a type SERIAL, PARALLEL,ONE_OF or GROUP of SERIAL/PARALLEL/ONE_OF.
ONE_OF mean only need acceptance.
*/

type User struct {
	Name       string // Identifier
	Order      int // For acceptance order.
	IsAccepted bool // Flagging
	Next       []*User

	// Couple is for checking 'ONE_OF' acceptance of GROUP.
	Couple []*User
}

type Workflow struct {
	Users []*User
}

func main() {
	fmt.Println("vim-go")
}i
