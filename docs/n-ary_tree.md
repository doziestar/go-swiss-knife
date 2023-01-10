# NTree

An n-ary tree is a tree data structure in which each node can have any number of children. This implementation of an n-ary tree uses a struct called `Node` to represent each node in the tree. The `Node` struct contains the data and a slice of child nodes. The `NTree` struct has a field `root` of type `Node`, which points to the root of the tree.

## Fields
- `root *Node`: A pointer to the root node of the tree.

## Methods

### Insert(data int, parent int)
- Adds a new node with the given data as a child of the node with the specified parent data.
- Time complexity: O(n) in the worst case

### Search(data int) bool
- Searches the tree for a node with the given data.
- Time complexity: O(n) in the worst case

### LevelOrder() []int
- Performs a level-order traversal of the tree, where it visits all the nodes of the same level before going on to the next level.
- Returns a slice containing the data of all the nodes in the order they were visited.
- Time complexity: O(n)

### Height() int
- Calculates the height of the tree, which is the maximum distance from the root to a leaf node.
- Time complexity: O(n)

### Size() int
- Returns the total number of nodes in the tree
- Time complexity: O(n)

### Delete(data int)
- deletes a node with the specified data from the tree.
- Time complexity: O(n) in the worst case

### Print()
- Traverses the tree in pre-order and prints the data of each node
- Time complexity: O(n)

Note: This implementation does not guarantee the uniquness of the value of nodes in the tree.

[//]: # ()