package main

import "fmt"

// This DFS code based on theory from : https://www.programiz.com/dsa/graph-dfs

// Vertex or node is an object that visited when doing searching.

type Vertex struct {
	AdjacentList []*Vertex
	IsVisited    bool
	Value        int
}

type Graph struct {
	Vertices int
	Vertexs  []*Vertex
	Visited  *[]int
	Stack    *[]int
}

func newGraph(totalVertices int) *Graph {
	graph := &Graph{
		Vertices: totalVertices,
	}
	graph.Vertexs = make([]*Vertex, graph.Vertices)
	graph.Visited = &[]int{}
	graph.Stack = &[]int{}

	return graph
}

func setEdge(graph *Graph, src, dest int) {
	if graph.Vertexs[src] == nil {
		graph.Vertexs[src] = &Vertex{
			AdjacentList: make([]*Vertex, graph.Vertices),
		}
	}
	if graph.Vertexs[dest] == nil {
		graph.Vertexs[dest] = &Vertex{
			AdjacentList: make([]*Vertex, graph.Vertices),
		}
	}
	graph.Vertexs[src].AdjacentList[dest] = graph.Vertexs[dest]
	graph.Vertexs[dest].AdjacentList[src] = graph.Vertexs[src]
}

func dfs(graph *Graph, src int) {
	fmt.Println(src)
	graph.Vertexs[src].IsVisited = true

	// add unvisited vertex to stack
	for i, vertex := range graph.Vertexs[src].AdjacentList {
		if vertex != nil && !vertex.IsVisited {
			*(graph).Visited = append(*(graph).Visited, i)
		}
	}

	// pop stack to visit
	var indexNextVertex int

	indexNextVertex, *graph.Visited = (*graph.Visited)[len(*graph.Visited)-1], (*graph.Visited)[:len(*graph.Visited)-1]
	dfs(graph, indexNextVertex)

	return
}

func main() {
	graph := newGraph(5)
	setEdge(graph, 0, 2)
	setEdge(graph, 2, 1)
	setEdge(graph, 2, 4)
	setEdge(graph, 1, 3)
	setEdge(graph, 0, 1)
	dfs(graph, 0)
}
