package main
import "fmt"

type stack struct{
	head int
	arr []int
}

func createaArray(capacity int) stack{
	arr := make([]int, capacity)
	s := stack{head:0,arr:arr}
	return s
}

func (s stack) pop() int{
	if(s.head == 0){
		return -1
	}else{
		s.head = s.head-1
		return s.arr[s.head]
	}
}

func (s stack) push(a int){
	s.arr[s.head] = a
	s.head = s.head+1
}


func main(){
	st := createaArray(5)
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	st.push(5)
	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.pop())
}