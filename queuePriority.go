package main

import "fmt"

type node struct {
	key      int
	priority int
	next     *node
}

type queue struct {
	rear  *node
	front *node
}

func makeQueue() *queue {
	q := &queue{}
	return q
}

func (q *queue) enqueue(a int, p int) {
	temp := &node{key: a, priority: p}
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

func (q *queue) getHighest() int {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return -1
	} else {
		p := q.front.priority
		a := q.front.key
		temp := q.front
		hp := q.front
		hpChild := q.front
		for temp != q.rear {
			if temp.next.priority > p {
				p = temp.next.priority
				a = temp.next.key
				hp = temp.next
				hpChild = temp
			}
			temp = temp.next
		}
		if hp == q.front {
			q.front = q.front.next
		} else if hp == q.rear {
			q.rear = hpChild
			q.rear.next = nil
		} else {
			hpChild.next = hp.next
			hp = nil
		}

		return a
	}
}

// func (q *queue) dequeue() int {
// 	if q.front == nil {
// 		fmt.Println("Queue is empty")
// 		return -1
// 	} else {
// 		a := q.front.key
// 		q.front = q.front.next
// 		return a
// 	}
// }

func main() {
	q := makeQueue()
	q.enqueue(1, 4)
	q.enqueue(2, 5)
	q.enqueue(3, 2)
	fmt.Println(q.getHighest())
	fmt.Println(q.getHighest())
	fmt.Println(q.getHighest())
	fmt.Println(q.getHighest())
}
