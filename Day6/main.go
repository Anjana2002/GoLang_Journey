package main

import (
	"bufio"
	"fmt"
	"os"
)

//receiver function
type Rectangle struct{
	width int
	height int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

//pointer receiver
type Counter struct {
	value int
}

func(c *Counter) Increment() {
	c.value++
}

//user input
func getInput() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello", name)

	var a,b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println("Sum: ", a+b)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence")
	text, _ := reader.ReadString('\n')
	fmt.Println("Your text: ", text)
}

func getInput2(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input, nil

}
func main() {
	rect := Rectangle{10, 5}
	fmt.Println("Area: ", rect.area())

	c := Counter{10}
	c.Increment()
	fmt.Println(c.value)

	getInput()

	reader := bufio.NewReader(os.Stdin)

	name, err := getInput2("Enter your name: ", reader)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Hello,", name)
}
