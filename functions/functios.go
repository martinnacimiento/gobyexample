package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// Multiple parameters of the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Variadic function
func plusPlusPlus(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	fmt.Println("1+2+3 =", plusPlus(1, 2, 3))
	fmt.Println("1+2+3+4+5 =", plusPlusPlus(1, 2, 3, 4, 5))
}
