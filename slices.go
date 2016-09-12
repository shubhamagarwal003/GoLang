package main
import "fmt"

func main(){
	s := make([]string, 3)
	fmt.Println(len(s))
	fmt.Println("s: ",s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s,"d")
	fmt.Println("s after: ",s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c after: ",c[1:3])
}