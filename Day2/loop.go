//loop
package main
import "fmt"

func loop() {
	
	x := 0
	for x < 5{
		fmt.Println("value of x is ", x)
		x++
	}
	for i :=0; i<5; i++ {
		fmt.Println("value of i is ", i)
	}
	fruits := []string{"apple", "mango", "grape", "orange"}
	for i:=0; i<len(fruits);i++{
		fmt.Println(fruits[i])
	}
	for index, value := range fruits{
		fmt.Printf("the position at index %v is %v\n", index, value)
	}
	for _,value:=range fruits{
		fmt.Printf("The value is %v\n", value)
	}

	for index, value := range fruits{
		if index == 1{
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("The value of pos %v is %v \n", index, value)
	}
}