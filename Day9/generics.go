//generics
package main

import "fmt"

//example 1
type Ordered interface {
	int | string | float64
}

func max[T Ordered] (a, b T) T {
	if a < b {
		return a
	}
	return b
}

//example 2
func Index[T comparable] (list []T, target T) int {
	for idx, value := range list {
		if value == target {
			return idx
		}
	}
	return -1
}
type Person struct{
	Age int
}

//example 3
type Box[T any] struct {
	items []T
}
func (b * Box[T])Add(item T) {
	b.items = append(b.items, item)
}
func (b *Box[T])Get()[]T{
	return b.items
}

func gen() {

	fmt.Println(max(1, 2))
	fmt.Println(max("Hello", "World"))
	fmt.Println(max(1.1, 2.1))

	listInt := []int{1, 20, 40}
	listFloat := []float64{1.20, 20.3, 40.2}
	listPerson := [] Person{{Age: 10}, {Age:25}}
	fmt.Println(Index(listInt, 20))
	fmt.Println(Index(listFloat, 40.2))
	fmt.Println(Index(listPerson, Person{Age: 25}))

	b := Box[any]{}
	b.Add(10)
	b.Add(20)
	b.Add("lang")
	fmt.Println(b.Get())

}