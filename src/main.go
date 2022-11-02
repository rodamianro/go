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
}
