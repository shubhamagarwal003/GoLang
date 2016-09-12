package main
import "fmt"
func main(){
	var a [5]int
	fmt.Println(a)
	for i := 0; i<5; i++ {
		fmt.Println(a[i])
	}
}