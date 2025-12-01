package main

import (
	"fmt"
	"math"
	"strings"
)

//functions
func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}
//return
func plusPlus(a, b, c int) int {
	return a + b + c
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

//multiple value return
func getInitials(n string)(string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _,v := range names{
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

var score = 90.1

//variadic function
func sum(nums ...int){
	fmt.Print(nums, " ")
	total := 0
	for _,num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sayGreeting("Alex")
	res := plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
	area := circleArea(5.5)
	fmt.Println("Area: ", area)

	cycleNames([]string{"John", "Sagar"}, sayGreeting)
	word1, word2 := getInitials("hello world")
	fmt.Println(word1, word2)
	word3, word4 := getInitials("hello")
	fmt.Println(word3, word4)

	//package scope
	sayGreeting("Rahul")
	for _,v := range points {
		fmt.Println(v)
	}
	showScore()

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}