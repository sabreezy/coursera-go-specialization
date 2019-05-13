package main

/*
Instructions: Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
Program should print “Not Found!” otherwise. The
program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

Submit your source code for the program, “findian.go”.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	us := ""
	fmt.Printf("Please enter a string to search for the name 'ian':\n")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		us = scanner.Text()
	}
	us = strings.ToLower(us)
	//us = strings.ReplaceAll(us, " ", "")
	if !(strings.HasPrefix(us, "i")) || !(strings.HasSuffix(us, "n")) || !(strings.Contains(us, "a")) {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Found!")
	}
}
