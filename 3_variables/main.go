package main

import "fmt"

func main() {
	var Name string = "Ayush" // We need to use this variable, we cannot leave it unused

	// here Go auto inter type. if we give the value direct
	var lang = "go-Lang"
	fmt.Println(Name)
	fmt.Println(lang)

	var name = "Ayush"
	var age = 19
	var is_student = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(is_student)

	// Shorthand syntax
	cpu := "Intel i5 4th Gen"
	fmt.Println(cpu)
}
