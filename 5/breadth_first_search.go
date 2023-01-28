package main

import "fmt"

// This BFS code based on theory from : https://www.programiz.com/dsa/graph-bfs

// Vertex or node is an object that visited when doing searching.

// If the result is not the same with other source code, check again because when looking at the graph, it's all about perspective.
type Vertex struct {
	AdjacentList []*Vertex
	IsVisited    bool
	Value        int
}

type Graph struct {
	Vertices int
	Vertexs  []*Vertex
	Stack    *[]int
}

func newGraph(totalVertices int) *Graph {
	graph := &Graph{
		Vertices: totalVertices,
	}
	graph.Vertexs = make([]*Vertex, graph.Vertices)
	graph.Stack = &[]int{}
	return graph
}

func setEdge(graph *Graph, src, dest int) {
	if graph.Vertexs[src] == nil {
		graph.Vertexs[src] = &Vertex{
			Value:        src,
			AdjacentList: make([]*Vertex, graph.Vertices),
		}
	}
	if graph.Vertexs[dest] == nil {
		graph.Vertexs[dest] = &Vertex{
			Value:        dest,
			AdjacentList: make([]*Vertex, graph.Vertices),
		}
	}
	graph.Vertexs[src].AdjacentList[dest] = graph.Vertexs[dest]
	graph.Vertexs[dest].AdjacentList[src] = graph.Vertexs[src]
}

func bfs(graph *Graph, src int) {

	fmt.Println(src)

	// flag current vertex to visited so this vertex can be skipped.
	graph.Vertexs[src].IsVisited = true

	for i, vertex := range graph.Vertexs[src].AdjacentList {
		if vertex != nil && !vertex.IsVisited {
			*graph.Stack = append(*graph.Stack, i)
		}
	}

	var nextVertex int
	for {
		if len(*graph.Stack) == 0 || graph.Stack == nil {
			break
		}

		// get next queue
		nextVertex, *graph.Stack = (*graph.Stack)[0], (*graph.Stack)[1:len(*graph.Stack)-1]
		if !graph.Vertexs[nextVertex].IsVisited {
			bfs(graph, nextVertex)
		}
	}
	return
}

func main() {
	graph := newGraph(5)
	setEdge(graph, 0, 2)
	setEdge(graph, 1, 3)
	setEdge(graph, 1, 2)
	setEdge(graph, 2, 4)
	bfs(graph, 0)
}
