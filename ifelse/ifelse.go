package main

import "fmt"

func main() {
	firstCondition := 7%2 == 0
	secondCondition := 8%4 == 0
	thirdCondition := 8%2 == 0

	if firstCondition {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if secondCondition {
		fmt.Println("8 is divisible by 4")
	}

	if thirdCondition || 7%2 == 0 {
		fmt.Println("either 8 is even or 7 is even")
	}

	if num := 88; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
