package main

import "fmt"

func variable_types() {
	//strings
	var firstName string = "Gopher"
	var lastName = "Developer"
	var middleName string

	fmt.Println(firstName, lastName, middleName)

	firstName = "Go"
	middleName = "Lang"
	fmt.Println(firstName, lastName, middleName)

	nickName := "Dev"
	fmt.Println(nickName)

	//int
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	//floats
	var heightOne float32 = 25.8
	var heightTwo float64 = 55555555555555555555555555555555.8 
	fmt.Println(heightOne, heightTwo)

	 //booleans
	 age := 45 
	 fmt.Println(age <= 50)
	 fmt.Println(age >= 50)
	 fmt.Println(age == 45)
	 fmt.Println(age != 50)
	 
	 if age<30{
		 fmt.Println("age is less than 30")
	 } else if age<40{
		 fmt.Println("age is less than 40")
	 } else {
		 fmt.Println("age is not less than 45")
	 }

}


