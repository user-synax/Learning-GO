package main

import "fmt"

func sum(nums ...int) (string, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return "Total Sum:", total
}

func main() {
	msg, total := sum(67, 4, 8, 5, 3, 2, 5, 32, 5)
	fmt.Println(msg, total)
}
