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

