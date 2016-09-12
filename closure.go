package main
import "fmt"

func fib(a int) int{
	if a == 1 {
		return 1
	} else if a == 2 {
		return 1
	} else {
		return fib(a-1) + fib(a-2)
	}
}

func main(){
	fmt.Println(fib(5))
}