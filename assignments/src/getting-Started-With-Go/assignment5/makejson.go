package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

func main() {
	m := make(map[string]string)
	fmt.Println("Please enter the name of the individual")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Please enter the address of the indivdual")
	scanner.Scan()
	address := scanner.Text()
	m[name] = address
	barr, _ := json.Marshal(m)
	fmt.Println("Outputting byte array from json marshall")
	fmt.Println(string(barr))
}
