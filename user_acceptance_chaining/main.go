package main

import (
	"user_acceptance_chain/entity"
)

func main() {
	user1 := entity.NewUser("John")
	user2 := entity.NewUser("Robert")
	user3 := entity.NewUser("Mike")
	user4 := entity.NewUser("Rudi")
	user5 := entity.NewUser("Tabuti")

	workflow := entity.NewWorkflow()
	workflow.
		InitUser(&user1).
		NextSerial().To(&user2).
		NextSequence().
		NextParallel().To(&user3).To(&user4).
		NextSequence().
		NextSerial().To(&user5)

	user2.PrintToLastConnectedUser()
}
