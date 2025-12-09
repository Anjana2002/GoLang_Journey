//printing and formatting strings
package main

import "fmt"

func printing() {
	age := 22 
    firstName := "Alex"
    lastName := "John"

    fmt.Print("hello")
    fmt.Print(" world\n")
    fmt.Println("My name is", firstName ,"and my age is", age)
    fmt.Printf("My name is %v and my age is %v\n", firstName, age)
    fmt.Printf("We are %q and %q\n", firstName , lastName)
    fmt.Printf("age is of type %T\n", age)
    fmt.Printf("you scored %f points \n", 25.5)
    fmt.Printf("you scored %0.1f points \n", 25.5555)
    var str = fmt.Sprintf("My age is %v and my name is %v\n", age, firstName)
    fmt.Println("The saved string is", str)
}