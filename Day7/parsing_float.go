//parsing float
package main

import (
	"fmt"
	"strconv"
)
func parsing_float() {
	input := "12.58"
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed float:", value)
}
