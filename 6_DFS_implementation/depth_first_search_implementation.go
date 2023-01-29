package main

import "fmt"

// This implementation is how to find a person name from another person.
// if that source person have a relation with wanted person, it will return have relation and it's vertex location.

type Person struct {
	Relations []*Person
	IsVisited bool
	Name      string
}

type Graph struct {
	TotalPerson int
	People      []*Person
	Stack       *[]int
}

func newGraph(totalPerson int) *Graph {
	graph := &Graph{
		TotalPerson: totalPerson,
		People:      make([]*Person, totalPerson),
	}
	graph.Stack = &[]int{}
	return graph
}

func addPerson(graph *Graph, identifier int, name string) {
	graph.People[identifier] = &Person{
		Relations: make([]*Person, graph.TotalPerson),
		Name:      name,
	}
}

func setRelation(graph *Graph, src, dest int) error {
	if graph.People[src] == nil {
		return fmt.Errorf("unknown source person")
	}
	if graph.People[dest] == nil {
		return fmt.Errorf("unknown destination person")
	}

	graph.People[src].Relations[dest] = graph.People[dest]
	graph.People[dest].Relations[src] = graph.People[src]
	return nil
}

func dfs(graph *Graph, name string, src int) (bool, int) {

	// flag current vertex to visited so this vertex can be skipped.
	graph.People[src].IsVisited = true
	if graph.People[src].Name == name {
		return true, src
	}

	for i, person := range graph.People[src].Relations {
		if person != nil && !person.IsVisited {
			*graph.Stack = append(*graph.Stack, i)
		}
	}

	var nextPerson int
	for {
		if len(*graph.Stack) == 0 || graph.Stack == nil {
			break
		}

		// pop the stack
		nextPerson, *graph.Stack = (*graph.Stack)[len(*graph.Stack)-1], (*graph.Stack)[:len(*graph.Stack)-1]
		if !graph.People[nextPerson].IsVisited {
			return dfs(graph, name, nextPerson)
		}
	}
	return false, -1
}

func main() {
	graph := newGraph(10)

	addPerson(graph, 0, "Yoga")
	addPerson(graph, 1, "Roger")
	addPerson(graph, 2, "John")
	addPerson(graph, 3, "Budi")
	addPerson(graph, 4, "Ardan")
	addPerson(graph, 5, "Raffa")
	addPerson(graph, 6, "Rose")
	addPerson(graph, 7, "Nayeon")
	addPerson(graph, 8, "Mina")
	addPerson(graph, 9, "Bae")

	setRelation(graph, 0, 3)
	setRelation(graph, 3, 4)
	setRelation(graph, 5, 0)
	setRelation(graph, 2, 4)
	setRelation(graph, 1, 5)

	setRelation(graph, 6, 7)
	setRelation(graph, 8, 9)
	setRelation(graph, 7, 9)

	if ok, location := dfs(graph, "Rose", 9); ok {
		fmt.Printf("is found in %d ", location)
	} else {
		fmt.Println("not found")
	}

}
