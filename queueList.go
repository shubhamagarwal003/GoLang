package main

import "fmt"

type node struct {
	key  int
	next *node
}

type queue struct {
	rear  *node
	front *node
}

func makeQueue() *queue {
	q := &queue{}
	return q
}

func (q *queue) enqueue(a int) {
	temp := &node{key: a}
	if q.rear == nil {
		q.rear = temp
		q.front = temp
	} else {
		q.rear.next = temp
		q.rear = temp
	}
	// q.rear.key = a
	// next := &node{}
	// q.rear.next = next
	// q.rear = q.rear.next
}

func (q *queue) dequeue() int {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return -1
	} else {
		a := q.front.key
		q.front = q.front.next
		return a
	}
}

func main() {
	q := makeQueue()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
