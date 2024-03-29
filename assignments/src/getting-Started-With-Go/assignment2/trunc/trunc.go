package main

import "fmt"

/*
Instructions:
Write a program which prompts the user to enter a floating point number
and prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/

func main() {
	var ui float64
	fmt.Printf("Please enter a floating point number:")
	fmt.Scan(&ui)
	fmt.Printf("Here is the integer version of the floating point value: %d\n", int64(ui))
}
