package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sortFirstQ(firstQ []int, ch1 chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("printing first Quarter", firstQ)
	sort.Ints(firstQ)
	fmt.Println("Printing sorted first Quarter", firstQ)
	ch1 <- firstQ
}

func sortSecondQ(secondQ []int, ch2 chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("printing second Quarter", secondQ)
	sort.Ints(secondQ)
	fmt.Println("Printing sorted second Quarter", secondQ)
	ch2 <- secondQ
}

func sortThirdQ(thirdQ []int, ch3 chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("printing third Quarter", thirdQ)
	sort.Ints(thirdQ)
	fmt.Println("Printing sorted third Quarter", thirdQ)
	ch3 <- thirdQ
}

func sortFourthQ(fourthQ []int, ch4 chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("printing fourth Quarter", fourthQ)
	sort.Ints(fourthQ)
	fmt.Println("Printing sorted fourth Quarter", fourthQ)
	ch4 <- fourthQ
}

func main() {
	usrArr := []int{}
	fmt.Println("Please enter a series of numbers separated by a space:")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	strArr := strings.Split(line, " ")
	for _, val := range strArr {
		val, _ := strconv.Atoi(val)
		usrArr = append(usrArr, val)
	}
	ch1 := make(chan []int)
	ch2 := make(chan []int)
	ch3 := make(chan []int)
	ch4 := make(chan []int)
	wg.Add(4)
	arr1 := usrArr[0:(len(usrArr) / 4)]
	arr2 := usrArr[(len(usrArr) / 4):(len(usrArr) / 2)]
	arr3 := usrArr[len(usrArr)/2 : (3 * (len(usrArr) / 4))]
	arr4 := usrArr[3*len(usrArr)/4:]
	go sortFirstQ(arr1, ch1, &wg)
	arr1 = <-ch1
	go sortSecondQ(arr2, ch2, &wg)
	arr2 = <-ch2
	go sortThirdQ(arr3, ch3, &wg)
	arr3 = <-ch3
	go sortFourthQ(arr4, ch4, &wg)
	arr4 = <-ch4
	wg.Wait()
	finalArr := []int{}
	finalArr = append(finalArr, arr1...)
	finalArr = append(finalArr, arr2...)
	finalArr = append(finalArr, arr3...)
	finalArr = append(finalArr, arr4...)
	sort.Ints(finalArr)
	fmt.Println("Printing Final Array:", finalArr)
}
