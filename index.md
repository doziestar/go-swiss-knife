# Linked List

A linked list is a data structure that consists of a sequence of elements, each of which contains a reference (or "link") to the next element in the sequence. This implementation of a linked list uses a struct to represent each element (or "node") in the list, and a pointer to the next node.

## Fields

- `head *Node`: A pointer to the first element in the list.
- `size int`: An integer to store the size of the list.

## Methods

### Append(data int)

- Appends an element to the end of the list.
- Time complexity: O(n)

### Insert(index int, data int)
- Inserts an element at a given position in the list.
- Time complexity: O(n)

### Delete(index int)
- Removes an element at a given position in the list.
- Time complexity: O(n)

### Print()
- Prints all the elements in the list
- Time complexity: O(n)

### Size() int
- Returns the number of elements in the list.
- Time complexity: O(1)

### Get(index int) (int, error)
- Retrieves the value at a given position in the list.
- Time complexity: O(n)

### Pop() (int, error)
- Removes and returns the last element in the list.
- Time complexity: O(n)

### Reverse()
- Reverses the order of the elements in the list.
- Time complexity: O(n)

### Clear()
- Removes all the elements from the list.
- Time complexity: O(1)

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

# Queue

A queue is a first-in, first-out (FIFO) data structure that provides two basic operations: `Enqueue` and `Dequeue`. This implementation of a queue uses a slice to store the elements of the queue, with the first element in the slice representing the front of the queue.

## Methods

### Enqueue(x int)
- Add an element to the back of the queue
- Time complexity: O(1)

### Dequeue() int
- Remove and return the front element from the queue
- Time complexity: O(n)

### Peek() int
- Returns the front element from the queue without removing it
- Time complexity: O(1)

### IsEmpty() bool
- Returns true if the queue is empty, false otherwise
- Time complexity: O(1)

### Size() int
- Returns the number of elements in the queue
- Time complexity: O(1)

Note: If queue is empty, Dequeue() and Peek() will cause a panic.

# Stack

A stack is a last-in, first-out (LIFO) data structure that provides two basic operations: `Push` and `Pop`. This implementation of a stack uses a slice to store the elements of the stack, with the last element in the slice representing the top of the stack.

## Methods

### Push(x int)
- Add an element to the top of the stack
- Time complexity: O(1)

### Pop() int
- Remove and return the top element from the stack
- Time complexity: O(1)

### Peek() int
- Returns the top element from the stack without removing it
- Time complexity: O(1)

### IsEmpty() bool
- Returns true if the stack is empty, false otherwise
- Time complexity: O(1)

### Size() int
- Returns the number of elements in the stack
- Time complexity: O(1)

Note: If stack is empty, Pop() and Peek() will cause a panic.

