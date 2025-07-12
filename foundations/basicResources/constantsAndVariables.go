package main

import (
	"fmt"
	mathLibAlias "math"
)

func main() {
	// Unusual
	const PI float64 = 3.1415
	var radius = 3.2 // Inferred typing

	// Shortened (and most common) way to declare variables
	area := PI * mathLibAlias.Pow(radius, 1)
	area = PI * mathLibAlias.Pow(radius, 2) // Changing value

	fmt.Println("Area:", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)

	fmt.Println("Another way to declare the variables", a, b, c, d)

	b1, b2 := true, "falseStr"

	fmt.Println("One-line multiple variable declaring:", b1, b2)
}
