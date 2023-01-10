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

