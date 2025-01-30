package main

import (
	"fmt"
	"math"
)

const nick string = "Tincho"

func main() {
	fmt.Println(nick)

	const n = 5000000 // 5M

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
