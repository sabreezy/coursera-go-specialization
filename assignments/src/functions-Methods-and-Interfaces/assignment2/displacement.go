package main

import "fmt"
import "math"

func genDisplaceFn(a, vo, so float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return (.5 * a * math.Pow(t, 2.0)) + (vo * t) + (so)
	}
	return fn
}

func main() {
	var a, vo, so, t float64
	fmt.Println("Please enter the value for acceleration")
	fmt.Scan(&a)
	fmt.Println("Please enter the value for initial velocity")
	fmt.Scan(&vo)
	fmt.Println("Please enter the value for initial displacement")
	fmt.Scan(&so)
	fmt.Println("Please enter a value of time")
	fmt.Scan(&t)
	fn := genDisplaceFn(a, vo, so)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
