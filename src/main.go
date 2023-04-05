/**
*Folder name where the file is saved
*The exception is main.go that is the root file*
**/
package main

// Package import
import (
	"fmt"
	"math"
)

// Main function
func main() {

	// Hellow world
	fmt.Println("Hello world") // Print the message in the console
	fmt.Println("")

	// Constants declaration
	const pi float64 = 3.14 // Firts form
	const pi2 = 3.1415      // Second form
	fmt.Println("Constans declaration")
	fmt.Println("\tFirst form (const pi float64 = 3.14):", pi)
	fmt.Println("\tSecond form (const pi2 = 3.1415):", pi2)
	fmt.Println("PD: The difference between the firts one and the second one is that the second one has implicit the type of constant")
	fmt.Println("")

	// Declaration of integer variables
	base := 12          // Firts form
	var height int = 14 // Second form
	var area int        // third form
	fmt.Println("Declaration of integer variables")
	fmt.Println("\tFirst form (base := 12):", base)
	fmt.Println("\tSecond form (var height int = 14):", height)
	fmt.Println("\tThird fomr (var area int):", area)
	fmt.Println("PD: the operator := is used when you want to create a variable without having declared it before")
	fmt.Println("")

	// Zero values
	var i int     // 0
	var f float64 // 0
	var s string  // Empty string
	var b bool    // false
	fmt.Println("Zero values")
	fmt.Println("\tInteger (var i int):", i)
	fmt.Println("\tFloat (var f float64):", f)
	fmt.Println("\tString (var s string):", s)
	fmt.Println("\tBoolean (var b bool):", b)
	fmt.Println("")

	// Operands
	fmt.Println("Operators")
	// Operamds
	x := 10
	y := 50
	fmt.Println("x :=", x)
	fmt.Println("y :=", y)

	// 1. Additon +
	result := x + y
	fmt.Println("\t1. Addition (result := x + y):", result)

	// 2. Subtraction -
	result = y - x
	fmt.Println("\t2. Subtraction (result = y - x):", result)

	// 3. Multiplication *
	result = x * y
	fmt.Println("\t3. Multiplication (result = x * y):", result)

	// 4. Division /
	result = y / x
	fmt.Println("\t4. Division (result = y / x):", result)

	// 5. Remainder %
	result = x % y
	fmt.Println("\t5. Remainder (result = x % y):", result)

	// 6. Incremental ++
	x++
	fmt.Println("\t6. Incremental (x++):", x)

	// 7. Decremental --
	x--
	fmt.Println("\t7. Decremental (x--):", x)

	// Exercises
	fmt.Println("Exersice")

	// 1. Square's area
	const squareBase = 10                                                                  // Declare the constant to save the square's base
	squareArea := squareBase * squareBase                                                  // Calculates the square's area
	fmt.Println("\t1. Square's area (squareArea := squareBase * squareBase):", squareArea) // Prints the result of calculates the square's area

	// 2. Rectangle's area
	const rectangleBase = 10                                                                                // Declare the constant to save the rectangle's base
	const rectangleHeight = 5                                                                               // Declare the constant to save the rectangle's height
	rectangleArea := rectangleBase * rectangleHeight                                                        // Calculates the rectangle's area
	fmt.Println("\t2. Rectangle's area (rectangleArea := rectangleBase * rectangleHeight):", rectangleArea) // Prints the result of calculates the rectangle's area

	// 3. Trapezoid
	// Definition: The trapezoid is a quadrilateral with at least one pair of parallel sides
	const firtsParallelSide int = 10                                                                                        // One parallel side
	const secondParralelSide int = 12                                                                                       // The other parallel side
	const trapezoidHeight int = 5                                                                                           // Trapezoid height
	var trapezoidArea int = ((firtsParallelSide + secondParralelSide) / 2) * trapezoidHeight                                // Calculates the trapezoid's area
	fmt.Println("\t3. Trapezoid's area (((firtsParallelSide + secondParralelSide) / 2) * trapezoidHeight):", trapezoidArea) // Prints the calculates result

	// 4. Circle's area
	const circleRadius = 5                                                                 // Circle's radius
	var ciclesArea = math.Pi * (math.Pow(circleRadius, 2))                                 // Calculates circule's area
	fmt.Println("\t4. Circle's area (math.Pi * (math.Pow(circleRadius, 2))):", ciclesArea) // Prints result calculates
	fmt.Println("")

	// Primitive values
	// Defining the value type improves performance
	fmt.Println("Primitive values")

	// Integer numbers
	fmt.Println(" Integer numbers")
	// int = OS depending (32 o 64)
	const integer int = 9223372036854775807
	fmt.Println("\tInteger that has OS depends (const integer int = 9223372036854775807):", integer)
	// int8 = 8 bits = -2^7 to 2^7 - 1 (-128 a 127)
	const integer8 int8 = 127
	fmt.Println("\tInteger that is equivalent to 8 bits (const integer8 int8 = 127):", integer8)
	// int16 = 16 bits = -2^15 to 2^15 - 1 (-32768 a 32767)
	const integer16 int16 = -32768
	fmt.Println("\tInteger that is equivalent to 16 bits (const integer16 int16 = -32768):", integer16)
	// int32 = 32 bits = -2^31 to 2^31 - 1 (-2147483648 a 2147483647)
	const integer32 int32 = 2147483647
	fmt.Println("\tInteger that is equivalent to 32 bits (const integer32 int32 = 2147483647):", integer32)
	// int64 = 64 bits = -2^63 to 2^63 - 1 (-9223372036854775808 a 9223372036854775807)
	const integer64 int64 = -9223372036854775808
	fmt.Println("\tInteger that is equivalent to 64 bits (const integer64 int64 = -9223372036854775808):", integer64)

	// Positive integer numbers
	fmt.Println(" Positive integer numbers")
	// uint = OS depending (32 o 64)
	const positiveInteger uint = 18446744073709551615
	fmt.Println("\tPositive integer that has OS dependes (const positiveInteger uint = 18446744073709551615):", positiveInteger)
	// uint8 = 8 bits = 0 to 2^8 -1 (0 to 255)
	const positiveInteger8 uint8 = 255
	fmt.Println("\tPositive integer that allows store numbers with a size of 8 bits (const positiveInteger8 uint8 = 255):", positiveInteger8)
	// uint16 = 16 bits = 0 to 2^16 - 1 (0 to 65535)
	const positiveInteger16 uint16 = 65535
	fmt.Println("\tPositive integer that allows store numbers with a size of 16 bits (const positiveInteger16 uint16 = 65535):", positiveInteger16)
	// uint32 = 32 bits = 0 to 2^32 - 1 (0 to 4294967295)
	const positiveInteger32 uint32 = 4294967295
	fmt.Println("\tPositive integer that allows store numbers with a size of 32 bits (const positiveInteger32 uint32 = 4294967295):", positiveInteger32)
	// uint64 = 64 bits = 0 to 2^64 - 1 (0 to 18446744073709551615)
	const positiveInteger64 uint64 = 18446744073709551615
	fmt.Println("\tPositive integer that allows store numbers with a size of 64 bits (const positveInteger64 uint64 = 18446744073709551615):", positiveInteger64)

	// Decimal numbers
	fmt.Println(" Decimal numbers")
	// float32 = 32 bits = +/- 1.18e^-38 to +/- 3.4e^38
	const decimal32 float32 = 91827389123891273897123897128397189237.123123123123
	fmt.Println("\tDecimal number that allows store numbers with size of 32 bits (const decimal32 float32 = 108310167974186771.14911748378442):", decimal32)
	// float64 = 64 bits = +/- 2.23e^-308 to +/- 1.8e^308
	const decimal64 float64 = 123123182318923923789127398127398712389712389712389712893712893712893789123789128907128923897128397123889123789123782337891237893789129123237899783212378989123237897891232372378931212312312123123231312312123123231231312312231123233333333333333333333333333333333333333333333333333333333333333333333333333333328.1
	fmt.Println(("\tDecimal number that allows store numbers with size of 64 bits (const decimal64 float64 = 1231231823...):"), decimal64)

	// Texts and booleans
	fmt.Println(" Texts and booleans")
	// Text
	const text string = ""
	fmt.Println("\tText (const text string = \"\"):", text)
	// Boolean
	const boolean bool = true
	fmt.Println("\tBoolean (const boolean bool = true):", boolean)

	// Complex numbers
	fmt.Println(" Complex numbers")
	// Complex64 = Real or imaginarie float32
	// Complex128 = Real or imaginarie float64
	const complexNumbers128 complex128 = 10 + 8i
	fmt.Println("\tComplex number (const complexNumber128 complex128 = 10 + 8i):", complexNumbers128)

	// fmt
	var helloMessage string = "Hello"
	var worldMessage string = "world"
	// Println
	fmt.Println(helloMessage, worldMessage)
	// Printf
	// In this case considers that % has a specific meaning
	fmt.Printf("%v %v\n", helloMessage, worldMessage)
	fmt.Println("For more verbs see: https://pkg.go.dev/fmt#hdr-Printing")
	// Sprintf
	message := fmt.Sprintf("%v %v", helloMessage, worldMessage)
	fmt.Println(message)
	printFunction(message, "world")
	printFunction(toString(returnValue(2)), "")
	var value1, value2 int = doubleReturn(2)
	var value3, _ int = doubleReturn(2)
	printFunction(toString(value1), toString(value2))
	printFunction(toString(value3))
}

func printFunction(message, subject string) {
	fmt.Println(message, "-", subject)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a * 2, a * 3
}

func toString(value any) string { return fmt.Sprint(value) }
