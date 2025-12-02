package main
import "fmt"

//pass by value
func updateName(x string) string {
	x = "John"
	return x
}
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

//pointers
func zeroval(ival int) {
    ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	//map
	menu := map[string]float64{
		"rice" : 20.6,
		"sugar" : 10.5,
		"coffee" : 14.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["sugar"])

	for k, v := range menu{
		fmt.Println(k, "-", v)
	}
	menu["rice"] = 100.5
	fmt.Println(menu)

	//empty map
	m := make(map[string]int)
	m["k1"] = 10
	m["k2"] = 20
	fmt.Println(m)

	//Non-pointer values - strings, ints, bools, floats, arrays, structs
	name := "Alex"
	name = updateName(name)
	fmt.Println(name)

	//pointer wrapper values- slices, maps, functions
	updateMenu(menu)
	fmt.Println(menu)


	//pointers
	i := 1
    fmt.Println("initial:", i)
    zeroval(i)
    fmt.Println("zeroval:", i)
	zeroptr(&i)
    fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}