package main
import "fmt"

type node struct{
	data int
	next *node
}

func insert (n *node, data int){
	newNode := node{data: data}
	for n.next != nil {
		n = n.next
	}
	n.next = &newNode
}

func printList(n *node){
	fmt.Println("Printing")
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func reverseList(n *node){
	prev := n
	current := n
	next := n.next
	for next != nil{
		prev = current
		current = next
		next = current.next
		current.next = prev
	}
	n.next = nil
	printList(current)
}

func merge(n1 *node, n2 *node) *node{
	result := &node{}
	if(n1 == nil){
		return n2
	}else if(n2 == nil){
		return n1
	}
	if(n1.data <= n2.data){
		result = n1
		result.next = merge(n1.next,n2)
	}else{
		result = n2
		result.next = merge(n1,n2.next)
	}
	return result
}

func sort(n1 node) *node{
	if(n1.next == nil){
		return &n1
	}else{
		temp := n1
		temp.next = nil
		return merge(&temp,sort(*n1.next))
	}
}

func detectLoop(n1 *node){
	slow := n1
	fast := n1
	for fast.next != nil && fast.next.next != nil{
		slow = slow.next
		fast = fast.next.next
		if(slow == fast){
			current := slow.next
			i := 1
			for current != slow{
				current = current.next
				i++
			}
			fmt.Println(i)
			ahead := n1
			back := n1
			for i > 0{
				ahead = ahead.next
				i--
			}
			fmt.Println(ahead.data)
			for ahead.next != back.next{
				ahead = ahead.next
				back = back.next
			}
			ahead.next = nil
			return
			// return true
		}
	}
	// return false
}

func main(){
	head := node{data: 1}
	insert(&head,5)
	insert(&head,2)
	insert(&head,-1)
	insert(&head,7)
	insert(&head,10)
	current := &head
	for current.next != nil{
		current = current.next
	}
	current.next = head.next
	// printList(&head)
	// m := sort(head)
	// printList(m)
	detectLoop(&head)
	printList(&head)
	// head2 := node{data: 2}
	// insert(&head2,3)
	// insert(&head2,5)
	// printList(&head)
	// // reverseList(&head)
	// m := merge(&head, &head2)
	// printList(m)
}