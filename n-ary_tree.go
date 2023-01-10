package main

import "fmt"

type NAryNode struct {
	data     int
	children []*NAryNode
}

type NTree struct {
	root *NAryNode
}

func (t *NTree) Insert(data int, parent int) {
	newNode := &NAryNode{data: data}
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	var queue []*NAryNode
	queue = append(queue, current)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.data == parent {
			node.children = append(node.children, newNode)
			return
		}
		queue = append(queue, node.children...)
	}
}

func (t *NTree) Search(data int) bool {
	if t.root == nil {
		return false
	}
	current := t.root
	var queue []*NAryNode
	queue = append(queue, current)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.data == data {
			return true
		}
		queue = append(queue, node.children...)
	}
	return false
}

// LevelOrder []int: This method performs a level-order traversal of the tree, where it visits all the nodes of the same level before going on to the next level. It returns a slice containing the data of all the nodes in the order they were visited.
func (t *NTree) LevelOrder() []int {
	if t.root == nil {
		return []int{}
	}
	var queue []*NAryNode
	var result []int
	queue = append(queue, t.root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.data)
		queue = append(queue, node.children...)
	}
	return result
}

// Height int: This method calculates the height of the tree, which is the maximum distance from the root to a leaf node.
func (t *NTree) Height() int {
	if t.root == nil {
		return 0
	}
	var queue []*NAryNode
	var height int
	queue = append(queue, t.root)
	for len(queue) != 0 {
		height++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			queue = append(queue, node.children...)
		}
		queue = queue[levelSize:]
	}
	return height
}

// Size int: This method returns the total number of nodes in the tree
func (t *NTree) Size() int {
	if t.root == nil {
		return 0
	}
	var queue []*NAryNode
	var size int
	queue = append(queue, t.root)
	for len(queue) != 0 {
		node := queue[0]
		size++
		queue = append(queue, node.children...)
		queue = queue[1:]
	}
	return size
}

// Print : This method traverses the tree in pre-order and prints the data of each node.
func (t *NTree) Print() {
	if t.root == nil {
		return
	}
	var stack []*NAryNode
	stack = append(stack, t.root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(node.data, " ")
		for i := len(node.children) - 1; i >= 0; i-- {
			stack = append(stack, node.children[i])
		}
	}
}
