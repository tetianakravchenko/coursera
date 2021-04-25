package main

import (
	"fmt"
	"time"
)

var count int = 0

func main() {
	/*
		Race conditions is a problem that might happen as a result
		of all possible interleavings. The problem itself: outcome
		of the program depends on interleavings, which are not deterministic
		(interleaving of instructions between concurrent tasks is unknown,
		and depends on OS, Go runtime).

		Race conditions can occur due to communication between tasks (or goroutines),
		for example: multiple tasks are trying to access and manipulate the same shared
		variable, in this way those multiple tasks are communicating with each other
		through this shared variable. Due to the uncertainty of the Goroutine scheduling
		mechanism, the results of the program are unpredictable.
	*/
	go inc()
	go inc()
	time.Sleep(time.Second)
	fmt.Println("Count = ", count)
}

func inc() {
	count++
}
