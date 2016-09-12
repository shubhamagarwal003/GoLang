package main
import "fmt"

type person struct{
	name string
	age int
}

func main(){
	s := person{"shubham",23}
	fmt.Println(s)
	s.name = "sfmk"
	fmt.Println(s)
}