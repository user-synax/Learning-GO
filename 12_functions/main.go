package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "GoLang", "JavaScript", "Python"
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	a := add(4, 2)
	fmt.Println(a)
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(getLanguages())
	fmt.Println(lang1)
	fmt.Println(lang2)
	fmt.Println(lang3)


	fn := func(a int) int {
		return 2
	}
	fn(2)


}
