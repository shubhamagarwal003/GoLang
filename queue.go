package main

import "fmt"

type queue struct {
	rear     int
	front    int
	arr      []int
	size     int
	capacity int
}

func makeQueue(capacity int) *queue {
	arr := make([]int, capacity)
	q := &queue{
		rear:     0,
		front:    0,
		arr:      arr,
		size:     0,
		capacity: capacity,
	}
	return q
}

func (q *queue) isFull() bool {
	return q.size == q.capacity
}
func (q *queue) enqueue(a int) {
	if q.isFull() {
		fmt.Println("Queue is full")
		return
	}
	q.arr[q.rear] = a
	q.rear = (q.rear + 1) % q.capacity
	q.size++
	// fmt.Println(q.size)
}

func (q *queue) dequeue() int {
	if q.size == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	a := q.arr[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return a
}

func main() {
	q := makeQueue(5)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
