package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	
	TwodArr := [2][2]int{{1,2},{3,4}}
	fmt.Println(TwodArr)

	Array := [...][2]int{{1,2}, {3,4}, {5,6}}
	for _, nested := range Array{
		for _, value := range nested{
			fmt.Println(value)
		}
	}
}