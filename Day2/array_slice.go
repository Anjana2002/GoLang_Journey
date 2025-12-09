//array and slicing

package main
import "fmt"

func array_slice() {
	
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
}