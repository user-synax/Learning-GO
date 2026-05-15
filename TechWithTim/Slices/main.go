package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[:]
	sl[0] = 100
	fmt.Println(sl, len(sl), cap(sl))

	slChars := []string{"Hello", "World"}
	for x := 0; x < 10; x++ {
		slChars = append(slChars, "Ayush")
		fmt.Println(slChars, len(slChars), cap(slChars))
	}
	fmt.Println(slChars)

	starts := []string{}
	for x := 0; x < 10; x++ {
		starts = append(starts, "*")
		fmt.Println(starts, len(starts), cap(starts))
	}
	fmt.Println(starts)

	tArr := []string{"Hello","World","Hii"}
	test(tArr)
	fmt.Println(tArr)
}

func test(arr []string) {
	arr[0] = "Changed This"
}