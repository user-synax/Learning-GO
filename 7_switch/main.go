package main

import (
	"fmt"
	"time"
)

func main() {
	// const password = 0000
	// switch password {
	// case 3423:
	// 	fmt.Println("Wrong")
	// case 1222:
	// 	fmt.Println("Wrong")
	// case 1000:
	// 	fmt.Println("Wrong")
	// case 1230:
	// 	fmt.Println("Wrong")
	// case 1234:
	// 	fmt.Println("Right")
	// default:
	// 	fmt.Println("Unable to Verify")
	// }

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("It's Monday")
	case time.Tuesday:
		fmt.Println("It's Tuesday")
	case time.Wednesday:
		fmt.Println("It's Wednesday")
	case time.Thursday:
		fmt.Println("It's Thursday")
	case time.Friday:
		fmt.Println("It's Friday")
	case time.Saturday:
		fmt.Println("It's Saturday")
	case time.Sunday:
		fmt.Println("It's Sunday")
	}

	// Function based [ simple ]

	WhoamI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("It's String")
		case bool:
			fmt.Println("It's Boolean")
		default:
			fmt.Println("Unable to identify Type")
		}
	}
	WhoamI(3.14)
}
