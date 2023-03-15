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
		NextSetParallel().To(&user2, &user3).
		NextSetSerial().To(&user4).
		NextSetParallel().To(&user5)

	user1.PrintToLastConnectedUser()
}
