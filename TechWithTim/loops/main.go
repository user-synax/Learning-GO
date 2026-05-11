package main

import "fmt"

func main() {
	// for a := 1; a < 5 + 1; a++ {
	// 	fmt.Println(a)
	// }

	name := "The fox came over the road and jump through it"
	// for n := 0; n < len(name); n++ {
	// 	fmt.Print(n)
	// }

	for index, char := range name {
		fmt.Println(index, string(char))
	}
}
