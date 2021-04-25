package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	capacity := 4
	c := make(chan []int, capacity)
	inputString := inputPrompt()
	inputInt, err := inputConvert(inputString)
	if err != nil {
		log.Fatal("Incorrect input!")
	}
	if len(inputInt) < 4 {
		log.Fatal("Please enter more integers.")
	}

	sl1, sl2 := splitSlice(inputInt)
	sl11, sl12 := splitSlice(sl1)
	sl21, sl22 := splitSlice(sl2)
	go sortSlice(sl11, c)
	go sortSlice(sl12, c)
	go sortSlice(sl21, c)
	go sortSlice(sl22, c)

	result := []int{}
	for i := 0; i < capacity; i++ {
		readCh := <-c
		result = append(result, readCh...)
	}
	sort.Ints(result)
	fmt.Printf("Final result: %v\n", result)
}

func sortSlice(sl []int, c chan []int) {
	fmt.Printf("Soting slice: %v\n", sl)
	sort.Ints(sl)
	c <- sl
}

func splitSlice(origin []int) ([]int, []int) {
	halfLen := len(origin) / 2
	return origin[0:halfLen], origin[halfLen:]
}

func inputPrompt() []string {
	fmt.Println("Enter integers separated by whitespace: ")
	scanner := bufio.NewReader(os.Stdin)
	s, _ := scanner.ReadString('\n')
	sTrim := strings.Trim(s, "\n")
	slStrings := strings.Split(sTrim, " ")
	return slStrings
}

func inputConvert(s []string) ([]int, error) {
	inputInt := make([]int, 0, len(s))
	for _, element := range s {
		intElem, err := strconv.Atoi(element)
		if err != err {
			return nil, err
		}
		inputInt = append(inputInt, intElem)
	}
	return inputInt, nil
}
