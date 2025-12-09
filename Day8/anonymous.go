//anonymous function -function without a name

package main
import "fmt"

func anonymous() {
	for i:= 0; i<5; i++ {
		func(j int) {
			fmt.Println("Printing", j, "from inside the anonymous function")
		}(i) //immediately invoked function expression
	}
}