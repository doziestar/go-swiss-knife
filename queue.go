package main

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		panic("Error: Queue is empty")
	}
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *Queue) Peek() int {
	if len(q.data) == 0 {
		panic("Error: Queue is empty")
	}
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}
