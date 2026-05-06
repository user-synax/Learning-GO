package main

import "fmt"

func main() {
	// if false {
	// 	fmt.Println("Yess")
	// } else {
	// 	fmt.Println("No")
	// }

	password := 1000
	if password == 1234 {
		fmt.Println("Correct Password")
	} else if password == 1000 {
		fmt.Println("Almost Here")
	} else {
		fmt.Println("InCorrect")
	}

	var role = "Admin"
	hasPermission := false

	if role == "Admin" && hasPermission {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}
}
