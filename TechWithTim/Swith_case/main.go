package main

import "fmt"

func main() {
	password := 1234

	switch password {
	case 1233 | 1234:
		fmt.Println("Incorrect")
	case 1234:
		fmt.Println("Correct")
	}
}
