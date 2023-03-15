package entity

import "fmt"

// Workflow is relation between users.
// In this context, workflow can be seen as graph where users is a nodes or vertexs.
type Workflow struct {
	Users     []*User
	lastOrder int
}

type Serial struct {
	workflow *Workflow
}

type Parallel struct {
	workflow *Workflow
}

func NewWorkflow() *Workflow {
	return &Workflow{
		lastOrder: 1,
	}
}

func (w *Workflow) InitUser(user *User) *Workflow {
	user.Order = w.lastOrder
	w.Users = append(w.Users, user)

	return w
}

func (w *Workflow) GetUsersFromOrder(n int) []*User {
	fmt.Println(n)
	var users []*User
	for _, user := range w.Users {
		if user.Order == n {
			fmt.Println(user.Name)
			users = append(users, user)
		}
	}
	return users
}

func (w *Workflow) NextSetSerial() *Serial {
	return &Serial{
		workflow: w,
	}
}

func (s *Serial) To(u *User) *Workflow {
	prevUsers := s.workflow.GetUsersFromOrder(s.workflow.lastOrder)

	s.workflow.lastOrder++
	u.Order = s.workflow.lastOrder
	for i := range prevUsers {
		prevUsers[i].Next = append(prevUsers[i].Next, u)
	}

	s.workflow.Users = append(s.workflow.Users, u)

	return s.workflow
}

func (w *Workflow) NextSetParallel() *Parallel {
	return &Parallel{
		workflow: w,
	}
}

func (p *Parallel) To(users ...*User) *Workflow {
	prevUsers := p.workflow.GetUsersFromOrder(p.workflow.lastOrder)

	p.workflow.lastOrder++
	for i := range users {
		users[i].Order = p.workflow.lastOrder
	}

	for i := range prevUsers {
		prevUsers[i].Next = append(prevUsers[i].Next, users...)
	}

	p.workflow.Users = append(p.workflow.Users, users...)

	return p.workflow
}
