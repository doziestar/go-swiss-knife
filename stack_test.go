package main

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}

	// Test Push and Peek
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Peek() != 3 {
		t.Errorf("Error: Peek method: expected 3, got %d", s.Peek())
	}

	// Test Pop
	val := s.Pop()
	if val != 3 {
		t.Errorf("Error: Pop method: expected 3, got %d", val)
	}
	if s.Size() != 2 {
		t.Errorf("Error: Pop method: expected size 2, got %d", s.Size())
	}

	// Test IsEmpty
	if !s.IsEmpty() {
		t.Errorf("Error: IsEmpty method: expected false, got true")
	}
	s.Pop()
	s.Pop()
	if s.IsEmpty() {
		t.Errorf("Error: IsEmpty method: expected true, got false")
	}

	//Test empty stack
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Error: Pop method: expected a panic")
		}
	}()
	s.Pop()
}
