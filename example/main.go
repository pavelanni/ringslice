package main

import (
	"fmt"

	"github.com/pavelanni/ringslice"
)

var months = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var monthsStr = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

func main() {
	rs := ringslice.New(&months)
	fmt.Println("index: ", 1, "value: ", rs.Value(1))
	fmt.Println("index: ", 12, "value: ", rs.Value(12))
	fmt.Println("index: ", 13, "value: ", rs.Value(13))
	fmt.Println("index: ", 144, "value: ", rs.Value(144))
	fmt.Println("index: ", 0, "value: ", rs.Value(0))
	fmt.Println("index: ", -3, "value: ", rs.Value(-3))
	fmt.Println("index: ", -13, "value: ", rs.Value(-13))

	rs.SetValue(1, 123)
	fmt.Println("Setting value for index 1 to 123")
	fmt.Println("index: ", 1, "value: ", rs.Value(1))

	rsStr := ringslice.New(&monthsStr)
	fmt.Println("index: ", 1, "value: ", rsStr.Value(1))
	fmt.Println("index: ", 12, "value: ", rsStr.Value(12))
	fmt.Println("index: ", 13, "value: ", rsStr.Value(13))
	fmt.Println("index: ", 144, "value: ", rsStr.Value(144))
	fmt.Println("index: ", 0, "value: ", rsStr.Value(0))
	fmt.Println("index: ", -3, "value: ", rsStr.Value(-3))
	fmt.Println("index: ", -13, "value: ", rsStr.Value(-13))

	rsStr.SetValue(1, "Blah")
	fmt.Println("Setting value for index 1 to Blah")
	fmt.Println("index: ", 1, "value: ", rsStr.Value(1))

}
