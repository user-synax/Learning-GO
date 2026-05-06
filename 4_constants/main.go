package main

import "fmt"

const age int = 19

func main() {
	const name string = "Ayush"
	fmt.Println(name)
	fmt.Println(age)

	// Multiple values in const
	const (
		age         int     = 19
		is_student  bool    = true
		value_of_PI float32 = 3.14
	)

	fmt.Println(age)
	fmt.Println(is_student)
	fmt.Println(value_of_PI)
}
