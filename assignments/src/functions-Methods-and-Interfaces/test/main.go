package main

import "fmt"

func Swap(sli []int, i, j int) {
	tmp := sli[i]
	sli[i] = sli[j]
	sli[j] = tmp
}

func BubbleSort(sli []int) {
	for i := len(sli) - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if sli[j-1] > sli[j] {
				Swap(sli, j-1, j)
			}
		}
	}
}

func main() {
	// input 10 nums
	sli := make([]int, 0)
	var x int
	for i := 0; i < 10; i++ {
		fmt.Println("Please input a number")
		fmt.Scan(&x)
		sli = append(sli, x)
	}

	BubbleSort(sli)
	fmt.Println(sli)

}
