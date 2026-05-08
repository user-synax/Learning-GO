package main

import (
	"fmt"
	"slices"
)

// import "slices"

// slices are dynamic arrays in GO
// Most used construct in GO
// + Useful methods
func main() {
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 1, 5)
	// nums = append(nums, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// Copy function
	// var numbers1 = make([]int, 0, 5)
	// numbers1 = append(numbers1, 1,2,3,4,5)
	// var numbers2 = make([]int, len(numbers1))
	// copy(numbers2, numbers1)
	// fmt.Println(numbers1)
	// fmt.Println(numbers2)

	// Slice operator

	// var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(nums[0:5])
	// fmt.Println(nums[0:])
	// fmt.Println(nums[:5])

	// Slice package
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2))
	fmt.Println(slices.Compare(nums1, nums2))
	fmt.Println(slices.Contains(nums1, 3))

	numbers := []int{1,2,7,7,3,2,5,8,5,2,4,5,7}
	slices.Sort(numbers)
	fmt.Println(numbers)

}
