package main
import "fmt"


// func test(s string){
// 	for i := 0; i<10;i++ {
// 		fmt.Println(s,":",i)
// 		message <- "abc"
// 	}
// }

func test1(done chan bool){
	fmt.Println("Test 1")
	done <- true
}

func test2(done chan bool){
	<- done
	fmt.Println("Test 2")

}

func main(){
	done := make(chan bool)
	go test2(done)
	go test1(done)
	var a int
	fmt.Scanln(&a)
}