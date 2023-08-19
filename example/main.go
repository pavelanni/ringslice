package main

import (
	"fmt"

	"github.com/pavelanni/ringslice"
)

var months = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func main() {
	rs := ringslice.NewInt(&months)
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

}
