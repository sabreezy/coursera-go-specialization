package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	userSlice := make([]int, 3) //Create a Slice of size 3

	var intToAdd int

	stop := false

	index := 0
	numActionsUser := 0

	for stop == false {

		fmt.Println("Please, enter a number to add:")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {

			if strings.ToLower(scanner.Text()) == "x" {
				fmt.Println("Program finished")
				stop = true
			}

			intToAdd, _ = strconv.Atoi(scanner.Text())
			numActionsUser++
		}

		if numActionsUser <= 3 {
			userSlice[index] = intToAdd
			index++
		} else {
			userSlice = append(userSlice, intToAdd)
		}

		sortedSlice := make([]int, len(userSlice))
		copy(sortedSlice, userSlice)

		sort.Ints(sortedSlice)
		fmt.Println("Sorted slice:", sortedSlice)

	}

}

/*
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
*/
