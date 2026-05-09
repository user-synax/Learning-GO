package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}

func main() {
	a := counter()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
