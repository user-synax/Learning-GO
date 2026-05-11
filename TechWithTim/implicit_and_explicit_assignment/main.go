package main

import "fmt"

func main() {
	var name string = "Ayush" // Explicit assignment we need to declare data type of a variable
	fmt.Println(name)

	username := "user-synax" // implicit assignment go automatically define the data-type
	fmt.Println(username)

	fmt.Printf("Username data-type = %T ", username)
}