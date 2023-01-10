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
