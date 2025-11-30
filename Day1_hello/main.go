package main
import (
    "fmt"
    "sort"
    "strings"
)

func main() {

    fmt.Println("Hello World!")
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

    //packages
    greeting := "hello there friends"
    fmt.Println(strings.Contains(greeting, "hello"))
    fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
    fmt.Println("orginal string: ", greeting)
    fmt.Println(strings.ToUpper(greeting))
    fmt.Println(strings.Index(greeting, "ll"))
    fmt.Println(strings.Split(greeting, " "))

    ages := []int {20, 16, 34, 23, 88, 45, 10}
    sort.Ints(ages)
    fmt.Println(ages)

    index := sort.SearchInts(ages, 88)
    fmt.Println(index)

    fruits := []string{ "orange", "grape", "jackfruit", "apple"}
    sort.Strings(fruits)
    fmt.Println(fruits)
    fmt.Println(sort.SearchStrings(fruits, "apple"))

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
    
    //printing and formatting strings
    age = 22 
    firstName = "Alex"
    lastName = "John"

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