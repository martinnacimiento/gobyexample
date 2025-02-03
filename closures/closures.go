package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		// This is an example of a closure. A closure is a function value that references variables from outside its body.
		i++ // i is declared in the outer function, but is accessible in the inner function
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
