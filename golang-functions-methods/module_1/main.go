package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputString := PromptInput()
	inputInt := ConvertInput(inputString)
	BubbleSort(inputInt)
	fmt.Printf("Sorted: %v\n", inputInt)
}

func PromptInput() []string {
	fmt.Println("Please enter up to 10 integers separated by whitespace: ")
	inputLimit := 10
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	sTrim := strings.Trim(s, "\n")
	slStrings := strings.Split(sTrim, " ")
	if len(slStrings) > inputLimit {
		fmt.Println("Input is too long")
		os.Exit(1)
	}
	return slStrings
}

func ConvertInput(input []string) []int {
	var slInt []int
	for _, v := range input {
		element, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Element '%v' can not be converted to int. Invalid input", v)
			os.Exit(1)
		}
		slInt = append(slInt, element)
	}
	return slInt
}

func BubbleSort(sl []int) {
	for i := 1; i < len(sl); i++ {
		for j := 0; j < len(sl)-1; j++ {
			if sl[j] > sl[j+1] {
				Swap(sl, j)
			}
		}
	}
}

func Swap(sl []int, index int) {
	tmp := sl[index]
	sl[index] = sl[index+1]
	sl[index+1] = tmp
}
