package main

import "fmt"

func main() {
	// 'var' can to infer the type, we'll create one variable 
	var name = "Martin"
	fmt.Println(name)
	
	// Declare multiple variables
	var firstname, lastname = "Martin", "Nacimiento"
	fmt.Println(firstname, lastname)

	// zero-valued
	var age int
	fmt.Println(age)

	// Shorthand for declaring and initialization
	profession := "Software Engineer"
	fmt.Println(profession)
}
