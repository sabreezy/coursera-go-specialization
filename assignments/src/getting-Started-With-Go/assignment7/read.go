package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Instructions: Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice,
and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

type person struct {
	fname string
	lname string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	people := make([]person, 0, 3)
	var fn string
	fmt.Println("Please indicate the name of the file: ")
	fmt.Scan(&fn)
	pwd, err := os.Getwd()
	fp := pwd + "/" + fn
	check(err)

	f, err := os.Open(fp)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var person person
		person.fname, person.lname = s[0], s[1]
		people = append(people, person)
	}

	f.Close()

	for _, v := range people {
		fmt.Println(v.fname, v.lname)
	}
}
