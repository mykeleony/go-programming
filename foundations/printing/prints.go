package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line.")

	fmt.Println("New")
	fmt.Println("line.")

	x := 3.1415
	// Does not work: fmt.Println("x value is" + x)

	stringx := fmt.Sprint(x)
	fmt.Println("x STRING value is " + stringx)
	fmt.Println("x FLOAT value is", x)
	fmt.Printf("x fprinted as float %.2f\n", x)
	fmt.Printf("x fprinted as string %s\n", x)

	a := 1
	b := 1.999999
	c := false
	d := "Example"

	fmt.Printf("a = %d, b = %f = %.1f, c = %t, %s\n", a, b, b, c, d)
	fmt.Printf("a = %v, b = %v = %.1v, c = %v, %v\n", a, b, b, c, d)
}
