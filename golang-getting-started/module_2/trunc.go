package main

import "fmt"

func main() {
	var floatNum float32
	fmt.Println("Enter a floating point number: ")
	fmt.Scan(&floatNum)
	fmt.Printf("Truncated value is %d\n", int(floatNum))
}
