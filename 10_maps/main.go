package main

import (
	"fmt"
	"maps"
)

// Maps is like object in js or dict in python, in go we called it maps
func main() {
	// Creating
	user := make(map[string]string)
	user["name"] = "Ayush"
	user["age"] = "19"
	fmt.Println(user)

	// Get specific element using key
	fmt.Println("Hi my name is:", user["name"])
	fmt.Println("I'm ", user["age"], "years old")

	// IMP: if key does not exist in map then it returns zero value
	fmt.Println(user["manoj"])

	ages := make(map[string]int)
	ages["Ayush"] = 19
	ages["Ritika"] = 18
	fmt.Println(ages["Ayush"])
	fmt.Println(ages["Manoj"]) // There is no key names 'manoj' in map so it will return 0

	fmt.Println(len(ages))
	fmt.Println(len(user))
	fmt.Println(user)
	delete(user, "age")
	fmt.Println(user)
	clear(user)
	fmt.Println(user)

	person := map[string]string{
		"name":  "Ayush",
		"age":   "19",
		"field": "Full stack web app developer",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	person["age"] = "21"
	fmt.Println(person["age"])
	fmt.Println(person["field"])
	delete(person, "field")
	fmt.Println(person)

	// Check if elemet avilable on map
	v, ok := person["name"]
	fmt.Println(v)
	if ok {
		fmt.Println("All ok")
	} else {
		fmt.Println("Not Ok")
	}

	m1 := map[string]int{"price": 90, "phone": 10}
	m2 := map[string]int{"price": 90, "phone": 10}
	// To check two maps if theu are same or not using inbuilt [ maps ] package
	fmt.Println(maps.Equal(m1,m2))

}
