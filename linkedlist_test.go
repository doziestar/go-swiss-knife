package main

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := LinkedList{}

	// Test Append method
	list.Append(1)
	list.Append(2)
	list.Append(3)
	if list.Size() != 3 {
		t.Errorf("Error: Append method: expected size 3, got %d", list.Size())
	}
	if val, _ := list.Get(0); val != 1 {
		t.Errorf("Error: Append method: expected 1 at index 0, got %d", val)
	}
	if val, _ := list.Get(1); val != 2 {
		t.Errorf("Error: Append method: expected 2 at index 1, got %d", val)
	}

	// Test Insert method
	list.Insert(1, 4)
	if list.Size() != 4 {
		t.Errorf("Error: Insert method: expected size 4, got %d", list.Size())
	}
	if val, _ := list.Get(1); val != 4 {
		t.Errorf("Error: Insert method: expected 4 at index 1, got %d", val)
	}

	// Test Delete method
	list.Delete(2)
	if list.Size() != 3 {
		t.Errorf("Error: Delete method: expected size 3, got %d", list.Size())
	}

	// Test Pop method
	val, _ := list.Pop()
	if val != 3 {
		t.Errorf("Error: Pop method: expected 3, got %d", val)
	}

	// Test Reverse method
	list.Append(5)
	list.Reverse()
	if val, _ := list.Get(0); val != 5 {
		t.Errorf("Error: Reverse method: expected 5 at index 0, got %d", val)
	}

	// Test Clear method
	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Error: Clear method: expected size 0, got %d", list.Size())
	}
}
