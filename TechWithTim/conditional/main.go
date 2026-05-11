package main

import "fmt"

func main() {
	age := 19
	if age >= 18 {
		fmt.Println("Grator than 18")
	}

	role := "Admin"
	login := 0
	if role == "Admin" {
		login += 999
		fmt.Println("Welcome, your login score:", login)
	}
}
