package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("Integers:", 1, 2, 1000)
	fmt.Println("Literal integer type is:", reflect.TypeOf(1000))

	// Unsigned
	const u1 uint8 = 1
	var b byte = 255

	fmt.Println("u1:", u1, "byte:", b)
	fmt.Println("Byte type is", reflect.TypeOf(b))

	// Signed
	MAX := math.MaxInt64
	fmt.Println("MAX:", MAX, "which type is:", reflect.TypeOf(MAX))

	var int32UnicodeValue rune = 'a'
	fmt.Println("Rune is", reflect.TypeOf(int32UnicodeValue))
	fmt.Println("Rune value is", int32UnicodeValue)

	// Real numbers
	var f float32 = 49.99
	var convertedF = float32(49.99)
	fmt.Println("Value", f, "has type", reflect.TypeOf(f))
	fmt.Println("Converted value", convertedF, "has type", reflect.TypeOf(convertedF))

	// Boolean
	boo := false
	fmt.Println("bool, which value is", boo, ", has type", reflect.TypeOf(boo))
	fmt.Println("Bool denial: ", !boo)

	// String
	str1 := "String"
	fmt.Println("String with value", str1, "has type:", reflect.TypeOf(str1), "and length", len(str1))

	str2 := `Multi
		line
		string`
	fmt.Println("Multi-line string:", str2, ", with length:", len(str2))

	// Char?
	char := 'a'
	fmt.Println("Char with value", char, "has type:", reflect.TypeOf(char))
	fmt.Printf("Printing it as char: %c", char)
}
