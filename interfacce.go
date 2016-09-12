package main
import "fmt"

type geometry interface{
	area() float64
}

type circle struct{
	radius float64
}

type square struct{
	side float64
}

func (s square) area() float64{
	return s.side*s.side
}

func (c circle) area() float64{
	return c.radius*c.radius*22/7
}

func measure(g geometry){
	fmt.Println(g.area())
}

func main(){
	s := square{side:5}
	c := circle{radius:2}
	measure(s)
	measure(c)
}