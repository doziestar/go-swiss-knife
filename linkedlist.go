package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	list.size++
}

func (list *LinkedList) Insert(index int, data int) {
	if index < 0 || index > list.size {
		fmt.Println("Error: Index out of range")
		return
	}
	newNode := &Node{data: data}
	if index == 0 {
		newNode.next = list.head
		list.head = newNode
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
	list.size++
}

func (list *LinkedList) Delete(index int) {
	if index < 0 || index >= list.size {
		fmt.Println("Error: Index out of range")
		return
	}
	if index == 0 {
		list.head = list.head.next
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	list.size--
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.size {
		return 0, fmt.Errorf("Error: Index out of range")
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data, nil
}

func (list *LinkedList) Pop() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Error: List is empty")
	}
	last := list.head
	for last.next != nil {
		last = last.next
	}
	data := last.data
	current := list.head
	for current.next != last {
		current = current.next
	}
	current.next = nil
	list.size--
	return data, nil
}

func (list *LinkedList) Reverse() {
	if list.head == nil || list.head.next == nil {
		return
	}
	var prev *Node
	current := list.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func (list *LinkedList) Clear() {
	list.head = nil
	list.size = 0
}

func (list *LinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
