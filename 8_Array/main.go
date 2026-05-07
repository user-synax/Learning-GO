package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums)

	// Array of booleans
	var bools [4]bool
	bools[0] = true

	fmt.Println(bools)

	// Array of strings
	var strs [4]string
	strs[2] = "Go Lang"
	fmt.Println(strs)

	// Array length
	fmt.Println("Total Length: ", len(nums))

	// Adding elements whil declarying...
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	TwoDArray := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(TwoDArray)
}
