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
	Queue    *[]int
}

func newGraph(totalVertices int) *Graph {
	graph := &Graph{
		Vertices: totalVertices,
	}
	graph.Vertexs = make([]*Vertex, graph.Vertices)
	graph.Queue = &[]int{}
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

// bfs using FIFO.
func bfs(graph *Graph, src int) {

	fmt.Println(src)

	// flag current vertex to visited so this vertex can be skipped.
	graph.Vertexs[src].IsVisited = true

	for i, vertex := range graph.Vertexs[src].AdjacentList {
		if vertex != nil && !vertex.IsVisited {
			*graph.Queue = append(*graph.Queue, i)
		}
	}

	var nextVertex int
	for {
		if len(*graph.Queue) == 0 || graph.Queue == nil {
			break
		}

		// get next queue
		nextVertex = dequeue(graph)
		if !graph.Vertexs[nextVertex].IsVisited {
			bfs(graph, nextVertex)
		}
	}
	return
}

func dequeue(graph *Graph) int {
	nextVertex := (*graph.Queue)[0]
	if len(*graph.Queue) == 1 {
		*graph.Queue = []int{}
	} else {
		*graph.Queue = (*graph.Queue)[1:]
	}

	return nextVertex
}

func main() {
	graph := newGraph(7)
	setEdge(graph, 0, 1)
	setEdge(graph, 0, 2)
	setEdge(graph, 0, 3)
	setEdge(graph, 2, 4)
	setEdge(graph, 2, 5)
	setEdge(graph, 3, 6)
	bfs(graph, 0)
}
