package main

import "fmt"

func main() {
	// Explicit type declaration
	var age int = 25
	var name string = "Adem"

	// Implicit type declaration (type inference)
	country := "Türkiye"

	// Multiple variable declaration
	var x, y, z int = 1, 2, 3

	// Output
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Country:", country)
	fmt.Println("Coordinates:", x, y, z)
}
