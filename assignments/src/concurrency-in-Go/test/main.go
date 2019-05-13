package main

import "fmt"

func printVal(arr []int, ch chan int) {
	for _, val := range arr {
		ch <- val
	}
}

func main() {
	initArr := []int{1, 23, 4, 5, 6, 7, 8, 10}
	ch := make(chan int)
	finalArr := []int{}
	go printVal(initArr, ch)
	for val := range ch {
		val = <-ch
		finalArr = append(finalArr, val)
	}
	fmt.Println(finalArr)
}
