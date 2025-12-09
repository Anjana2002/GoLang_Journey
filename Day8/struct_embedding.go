//struct embedding
package main

import "fmt"

//base struct
type Vehicle struct {
	brand string
}

//method of vehicle
func (v Vehicle) info() string {
	return "Brand: " + v.brand
}

//conatiner struct with embedding
type Car struct {
	Vehicle
	model string
}

//method of car
func (c Car) details() string {
	return c.info() + ", Model: " +c.model
}

//interface
type Describer interface {
	details() string
}

func embedding() {
	c := Car{
		Vehicle: Vehicle{brand: "tesla"},
		model :"model S",
	}

	fmt.Println(c.brand)
	fmt.Println(c.info())
	fmt.Println(c.details())

	var d Describer = c
	fmt.Println("Describer: ", d.details())
}