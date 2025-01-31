package main

import "fmt"

func main() {
	i := 1

	// Single condition
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Classic for loop, with initial/condition/after
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	// Do this N times
	for i := range 5 {
		fmt.Println("range: ", i)
	}

	// Without condition, infinite loop
	for {
		fmt.Println("Infinite loop")
		break
	}

	// Continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
