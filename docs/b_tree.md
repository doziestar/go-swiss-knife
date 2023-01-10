# Binary Tree

A binary tree is a tree data structure in which each node has at most two children, which are referred to as the left child and the right child. This implementation of a binary tree uses a struct called `Node` to represent each node in the tree. The `Node` struct contains the data and references to the left and right child nodes.

## Fields
- `root *Node`: A pointer to the root node of the tree.

## Methods

### Insert(data int)
- Adds a new node with the given data to the tree.
- Time complexity: O(n) in the worst case, when the tree is not balanced.

### Search(data int) bool
- Searches the tree for a node with the given data.
- Time complexity: O(n) in the worst case, when the tree is not balanced.

Note: This implementation does not handle duplicate values and the tree is not guaranteed to be balanced, which can lead to worst-case time complexities of O(n) on some operations.
