package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["age"] = 25
	m["height"] = 182

	fmt.Println(m)

	weight := m["weight"]
	fmt.Println(weight) // zero value

	_, ok := m["weight"] // check if key exists
	fmt.Println(ok)

	fmt.Println("Number of key-value pairs: ", len(m))

	delete(m, "height") // delete key-value pair
	fmt.Println(m)

	clear(m) // clear map
	fmt.Println(m)

	// map literal
	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(n)
	fmt.Println(n2)
	if maps.Equal(n, n2) {
		fmt.Println("Maps are equal")
	}
}
