package main

type BNode struct {
	data  int
	left  *BNode
	right *BNode
}

type BinaryTree struct {
	root *BNode
}

func (t *BinaryTree) Insert(data int) {
	newNode := &BNode{data: data}
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	for {
		if data <= current.data {
			if current.left == nil {
				current.left = newNode
				break
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				break
			}
			current = current.right
		}
	}
}

func (t *BinaryTree) Search(data int) bool {
	current := t.root
	for current != nil {
		if data == current.data {
			return true
		}
		if data < current.data {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

func (t *BinaryTree) Delete(data int) {
	if t.root == nil {
		return
	}
	var parent *BNode
	current := t.root
	for current != nil && current.data != data {
		parent = current
		if data < current.data {
			current = current.left
		} else {
			current = current.right
		}
	}
	if current == nil {
		return
	}
	if current.left == nil && current.right == nil {
		if parent == nil {
			t.root = nil
			return
		}
		if parent.left == current {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}
	if current.left == nil {
		if parent == nil {
			t.root = current.right
			return
		}
		if parent.left == current {
			parent.left = current.right
		} else {
			parent.right = current.right
		}
		return
	}
	if current.right == nil {
		if parent == nil {
			t.root = current.left
			return
		}
		if parent.left == current {
			parent.left = current.left
		} else {
			parent.right = current.left
		}
		return
	}
	var successor *BNode
	var successorParent *BNode
	successor = current.right
	for successor.left != nil {
		successorParent = successor
		successor = successor.left
	}
	if successorParent != nil {
		successorParent.left = successor.right
	} else {
		current.right = successor.right
	}
	current.data = successor.data
}

func (t *BinaryTree) InOrderTraversal() {
	t.inOrderTraversal(t.root)
}

func (t *BinaryTree) inOrderTraversal(node *BNode) {
	if node == nil {
		return
	}
	t.inOrderTraversal(node.left)
	print(node.data, " ")
	t.inOrderTraversal(node.right)
}

func (t *BinaryTree) PreOrderTraversal() {
	t.preOrderTraversal(t.root)
}

func (t *BinaryTree) preOrderTraversal(node *BNode) {
	if node == nil {
		return
	}
	print(node.data, " ")
	t.preOrderTraversal(node.left)
	t.preOrderTraversal(node.right)
}

// PostOrderTraversal print the tree in post-order traversal
func (t *BinaryTree) PostOrderTraversal() {
	t.postOrderTraversal(t.root)
}

func (t *BinaryTree) postOrderTraversal(node *BNode) {
	if node == nil {
		return
	}
	t.postOrderTraversal(node.left)
	t.postOrderTraversal(node.right)
	print(node.data, " ")
}
