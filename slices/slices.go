package main

import (
	"fmt"
	"slices"
)

// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

/**
Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
An uninitialized slice is nil, which has a length and capacity of 0.
*/

func main() {
	var a [5]string // This is an array of strings
	var s []string  // This is a slice of strings

	/**
	What is the difference between an array and a slice?
	- An array is a fixed length sequence of elements of a single type.
	- A slice is a dynamically-sized, flexible view into the elements of an array.
	- Memory is allocated for the array, but not for the slice. Slice is a reference.
	*/

	fmt.Println("uninit: ", s, &s, s == nil, len(s) == 0)
	fmt.Println("emp: ", a, &a, &a == nil, len(a) == 5)

	var notes = make([]int, 3)
	fmt.Println("emp:", notes, "len:", len(notes), "cap:", cap(notes))

	// set and get
	notes[0] = 1
	notes[1] = 2
	notes[2] = 3

	fmt.Println("set:", notes)
	fmt.Println("get:", notes[2])

	// append
	notes = append(notes, 4)
	fmt.Println("append:", notes)

	// Slices can be copied
	copiedNotes := make([]int, len(notes))
	copy(copiedNotes, notes)

	fmt.Println("copied:", copiedNotes)

	// Slices support a "slice" operator
	slice1 := notes[1:3] // slice up to but excluding the last index, i.e. [1, 3)
	fmt.Println("slice1:", slice1)

	slice2 := []string{"a", "b", "c", "d"}
	fmt.Println("slice2:", slice2)

	slices.Reverse(slice1)
	min := slices.Min(slice1)
	fmt.Println("reversed slice1:", slice1)
	fmt.Println("min:", min)

	// Slices can be composed into multi-dimensional data structures
	twoD := make([][]int, 3)
	fmt.Println("twoD:", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
