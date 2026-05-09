package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	// Normal method to iterate over a slice
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i]+1)
	// }

	// using range
	sum := 0
	for index, num := range nums {
		sum += num
		fmt.Println("Index",index, "Value",num)
	}
	fmt.Println("Total sum:", sum)

	m := map[string]string{
		"fname":"Jhon",
		"lname":"Doe",
	}

	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	// Using range on string

	for i, c := range "GoLang"{
		fmt.Println(i, string(c))
	}

}
