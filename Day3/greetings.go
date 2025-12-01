//package scope
package main
import "fmt"

var points = []int{10, 20, 30, 40}
func sayHello(n string){
	fmt.Println("hello", n)
}

func showScore() {
	fmt.Println("Score: ", score)
}