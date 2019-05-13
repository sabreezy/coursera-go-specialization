package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var x int

/*
This programs shows how to raise a race condition.
There is a shared variable called 'x' and two goroutines.
Each goroutine make an increment of x, so the output should be:
	Value 1
	Value 2

But there is not like this. The output is
	Value 2
	Value

You can see why in the comments in the code below
*/
func main() {
	x = 0

	//This goroutine emulates a long operation after updating the shared variable
	go func(i *int) {
		//the shared variable here is set to 1
		*i++
		time.Sleep(500)
		//Here, the variable should be 1, but is not, is set to 2 by the second goroutine
		fmt.Println("Value", *i)
	}(&x)

	//Since this goroutine is executed after the increment in the first goroutine but before the sleep(500)
	go func(i *int) {
		//The variable here is 1, so it is set to 2
		*i++
		//Prints 2
		fmt.Println("Value", *i)
	}(&x)

	bufio.NewReader(os.Stdin).ReadBytes('\n')

}
