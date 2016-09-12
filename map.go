package main
import "fmt"

func main() {
	m := make(map[string]int)
	m["d1"] = 1
	m["d2"] = 2
	fmt.Println(m)
	delete(m,"d2")
	fmt.Println(m)
	n := map[string]int{"d1": 1,"d2": 2}
	fmt.Println(n)
	a,b := n["d1"]
	fmt.Println(a,b)
}