package main

import (
	"fmt"
	"os"
	"strconv"
)

/*Test the program by running it twice and testing it with a different sequence of integers each time.

The first test sequence of integers should be all positive numbers and the second test should have at least one negative number.
Give 3 points if the program works correctly for one test sequence,
and give 3 more points if the program works correctly for the second test sequence.

Examine the code.
If the code contains a function called BubbleSort() which takes a slice of integers as an argument,
then give another 2 points.
If the code contains a function called Swap() function which takes two arguments,
a slice of integers and an index value i, then give another 2 points.
============================================================================================================================
Write a Bubble Sort program in Go.
The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program,
you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.
A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing,
but it should swap the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

func convCLIints(CLIargs []string) []int {
	usrVals := []int{}
	for i := 1; i < len(CLIargs); i++ {
		val, _ := strconv.Atoi(CLIargs[i])
		usrVals = append(usrVals, val)
	}
	return usrVals
}

func swap(v1, v2 *int) {
	var temp = *v1
	*v1 = *v2
	*v2 = temp
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func main() {
	posArr := []int{5, 7, 2, 4, 7, 1, 3, 5, 7, 8, 9}
	posArrNeg := []int{6, 4, 3, 25, 6, 7, 8, -1, 2}
	usrArr := []int{}
	fmt.Println("Please enter a sequence of numbers to run through the bubble sort!")
	usrArr = convCLIints(os.Args[1:]) //start at index 1 to ignore program name as part of os.Args
	fmt.Println(usrArr)
	fmt.Println(posArr)
	fmt.Println(posArrNeg)
	bubbleSort(usrArr)
	bubbleSort(posArr)
	bubbleSort(posArrNeg)
	fmt.Println(usrArr, posArr, posArrNeg)
}
