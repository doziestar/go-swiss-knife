package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func (g *Graph) AddVertex(v int) {
	g.adjList[v] = []int{}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.adjList[v1] = append(g.adjList[v1], v2)
	g.adjList[v2] = append(g.adjList[v2], v1)
}

func (g *Graph) Print() {
	for v, adj := range g.adjList {
		fmt.Printf("%d: ", v)
		for _, n := range adj {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}
}

// DFS DFS(start int): This method performs a Depth-First Search (DFS) traversal of the graph, starting from the given vertex. It uses a stack to keep track of the vertices to visit.
func (g *Graph) DFS(start int) {
	visited := map[int]bool{}
	var stack []int
	stack = append(stack, start)
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[current] {
			visited[current] = true
			fmt.Print(current, " ")
			for _, vertex := range g.adjList[current] {
				stack = append(stack, vertex)
			}
		}
	}
}

// BFS BFS(start int): This method performs a Breadth-First Search (BFS) traversal of the graph, starting from the given vertex. It uses a queue to keep track of the vertices to visit.
func (g *Graph) BFS(start int) {
	visited := map[int]bool{}
	var queue []int
	queue = append(queue, start)
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		if !visited[current] {
			visited[current] = true
			fmt.Print(current, " ")
			for _, vertex := range g.adjList[current] {
				queue = append(queue, vertex)
			}
		}
	}
}
