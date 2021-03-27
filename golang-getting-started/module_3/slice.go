package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sl := make([]int, 3)
	fmt.Println("Enter integers (X to exit): ")
	for {
		var v string
		fmt.Scan(&v)
		if v == "X" {
			break
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("can't convert input to int")
		}
		sl = append(sl, i)
		sort.Ints(sl)
		fmt.Println(sl)
	}
}
