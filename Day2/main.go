package main
import "fmt"

func main() {
	//array and slicing
	var ages [3]int = [3]int {10, 20, 30}
	var age = [3]int {10, 20, 30}
	names := [7]string{"a", "b", "c", "d", "g", "h", "i"}
	names[1] = "f"

	fmt.Println(ages, len(ages))
	fmt.Println(age)
	fmt.Println(names, len(names))

	var scores = []int{100, 90, 80}
	scores[2] = 25
	scores = append(scores, 70)
	fmt.Println(scores)
	
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)
	
	//loops
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