package main

import "testing"

func TestQueue(t *testing.T) {
	q := Queue{}

	// Test Enqueue and Peek
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Peek() != 1 {
		t.Errorf("Error: Peek method: expected 1, got %d", q.Peek())
	}

	// Test Dequeue
	val := q.Dequeue()
	if val != 1 {
		t.Errorf("Error: Dequeue method: expected 1, got %d", val)
	}
	if q.Size() != 2 {
		t.Errorf("Error: Dequeue method: expected size 2, got %d", q.Size())
	}

	// Test IsEmpty
	if !q.IsEmpty() {
		t.Errorf("Error: IsEmpty method: expected false, got true")
	}
	q.Dequeue()
	q.Dequeue()
	if q.IsEmpty() {
		t.Errorf("Error: IsEmpty method: expected true, got false")
	}

	//Test empty queue
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Error: Dequeue method: expected a panic")
		}
	}()
	q.Dequeue()
}
