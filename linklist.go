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

func main(){
	head := node{data: 1}
	printList(&head)
	insert(&head,2)
	printList(&head)
	insert(&head,3)
	printList(&head)
	insert(&head,4)
	printList(&head)

}