package main

import "fmt"

func main() {
	// uint - Unsigned Integer [ uint-8, uint-16, uint-32, uint-64]
	// int - Signed Integer [ represent negitive values too ] [ int-8, int-16, int-32, int-64]
	// float - decimal / floating point values [ float-32, float-64 ] ex: 2.3.14
	// byte = int-8
	// rune = int-32
	// bool = true or false
	// string = inside everything double qoutes [ in go single qoutes are not avilable ]
	// nil = Null or undefined

	var x string = "Hello World"
	fmt.Println(x)

	var y uint8  = 100
	fmt.Println(y)
}