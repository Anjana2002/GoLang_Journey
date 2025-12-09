package main
import (
	"fmt"
	"sort"
    "strings"
)
 func packages() {

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
 }