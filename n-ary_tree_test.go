package main

import (
	"reflect"
	"testing"
)

func TestNTree(t *testing.T) {
	tree := NTree{}

	// Test Insert
	tree.Insert(1, 0)
	tree.Insert(2, 1)
	tree.Insert(3, 1)
	tree.Insert(4, 2)
	tree.Insert(5, 2)
	tree.Insert(6, 3)
	tree.Insert(7, 3)

	// Test Search
	if !tree.Search(4) {
		t.Errorf("Error: Search method: expected true, got false")
	}
	if tree.Search(8) {
		t.Errorf("Error: Search method: expected false, got true")
	}

	// Test LevelOrder
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(tree.LevelOrder(), expected) {
		t.Errorf("Error: LevelOrder method: expected %v, got %v", expected, tree.LevelOrder())
	}

	// Test Height
	if tree.Height() != 4 {
		t.Errorf("Error: Height method: expected 4, got %d", tree.Height())
	}

	// Test Size
	if tree.Size() != 7 {
		t.Errorf("Error: Size method: expected 7, got %d", tree.Size())
	}
}
