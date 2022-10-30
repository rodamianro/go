/**
*Folder name where the file is saved
*The exception is main.go that is the root file*
**/
package main

// Package import
import (
	"fmt"
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
	fmt.Println("First form (const pi float64 = 3.14):", pi)
	fmt.Println("Second form (const pi2 = 3.1415):", pi2)
	fmt.Println("PD: The difference between the firts one and the second one is that the second one has implicit the type of constant")
	fmt.Println("")

	// Declaration of integer variables
	base := 12          // Firts form
	var height int = 14 // Second form
	var area int        // third form
	fmt.Println("Declaration of integer variables")
	fmt.Println("First form (base := 12):", base)
	fmt.Println("Second form (var height int = 14):", height)
	fmt.Println("Third fomr (var area int):", area)
	fmt.Println("PD: the operator := is used when you want to create a variable without having declared it before")
	fmt.Println("")

	// Zero values
	var i int     // 0
	var f float64 // 0
	var s string  // Empty string
	var b bool    // false
	fmt.Println("Zero values")
	fmt.Println("Integer (var i int):", i)
	fmt.Println("Float (var f float64):", f)
	fmt.Println("String (var s string):", s)
	fmt.Println("Boolean (var b bool):", b)
	fmt.Println("")

	// Area square
	const squareBase = 10                   // Declare the constant to save the square's base
	squareArea := squareBase * squareBase   // Calculates the square's area
	fmt.Println("Square area:", squareArea) // Prints the result of calculates the square's area
}
