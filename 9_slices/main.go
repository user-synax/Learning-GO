package main

import "fmt"

// import "slices"

// slices are dynamic arrays in GO
// Most used construct in GO
// + Useful methods
func main() {
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 1, 5)
	nums = append(nums, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9)
	fmt.Println(nums)
	fmt.Println(cap(nums))
}
