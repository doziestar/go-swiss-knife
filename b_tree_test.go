package main

import "testing"

func TestBinaryTree(t *testing.T) {
	tree := BinaryTree{}

	// Test Insert
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(8)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(9)

	// Test Search
	if !tree.Search(2) {
		t.Errorf("Error: Search method: expected true, got false")
	}
	if tree.Search(6) {
		t.Errorf("Error: Search method: expected false, got true")
	}
}
