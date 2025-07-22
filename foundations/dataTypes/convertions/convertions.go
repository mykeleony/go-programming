package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println("Integer to float:", x/float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println("Final grade converted to int:", finalGrade)

	fmt.Println("Careful concatenating strings, as string(123), for example, returns the 123 ASCII" +
		" corresponding char instead of the number 123 as string: " + string(123))

	// Integer to String
	fmt.Println("Integer to string:", strconv.Itoa(123))

	// String to Integer
	num, _ := strconv.Atoi("123")
	fmt.Println("String to int: ", num)

	// Boolean to Integer
	b, _ := strconv.ParseBool("-1")

	if b {
		fmt.Println("It is true!")
	} else {
		fmt.Println("It is not true!")
	}
}
