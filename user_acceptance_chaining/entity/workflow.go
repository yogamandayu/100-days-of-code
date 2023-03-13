package entity

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
	user.OrderTurn = w.lastOrder
	user.IsAccepted = true
	w.Users = append(w.Users, user)
	return w
}

func (w *Workflow) NextSerial() *Serial {
	return &Serial{
		workflow: w,
	}
}

func (s *Serial) To(user *User) *Serial {
	s.workflow.lastOrder++
	user.OrderTurn = s.workflow.lastOrder
	s.workflow.Users = append(s.workflow.Users, user)
	return s
}

func (s *Serial) NextSequence() *Workflow {
	return s.workflow
}

func (s *Serial) CanRepresentedBy(user *User) *Serial {
	return s
}

func (w *Workflow) NextParallel() *Parallel {
	w.lastOrder++
	return &Parallel{
		workflow: w,
	}
}

func (p *Parallel) NextSequence() *Workflow {
	return p.workflow
}

func (p *Parallel) To(user *User) *Parallel {
	user.OrderTurn = p.workflow.lastOrder
	p.workflow.Users = append(p.workflow.Users, user)
	return p
}

func (p *Parallel) CanRepresentedBy(user *User) *Parallel {
	return p
}
