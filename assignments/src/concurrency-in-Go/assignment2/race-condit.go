package main

import "fmt"

func changeVal(x *int) {
	*x = 5
	return
}
func getValue() int {
	var v int
	go changeVal(&v)
	return v
}

func main() {
	for {
		fmt.Println(getValue())
	}
}
