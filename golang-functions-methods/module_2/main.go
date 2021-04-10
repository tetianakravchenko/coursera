package main

import (
	"fmt"
	"math"
)

func main() {
	var a, v0, s0, t float64
	fmt.Printf("Enter acceleration value: ")
	fmt.Scan(&a)
	fmt.Printf("Enter initial velocity value: ")
	fmt.Scan(&v0)
	fmt.Printf("Enter initial displacement value: ")
	fmt.Scan(&s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Enter t: ")
	fmt.Scan(&t)
	fmt.Println(fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fn
}
