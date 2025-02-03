package main

import "fmt"

func vals() (int, int, string) {
	return 3, 7, "foo"
}

func main() {
	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	_, d, _ := vals()
	fmt.Println(d)
}
