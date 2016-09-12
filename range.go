package main
import "fmt"

func main(){
	nums := []int{1,2,3}
	var sum int
	for _,num := range(nums){
		sum = sum+num
	}
	fmt.Println(sum)
}