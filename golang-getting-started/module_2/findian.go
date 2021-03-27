package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	sTrim := strings.ReplaceAll(s, "\n", "")
	sLower := strings.ToLower(sTrim)
	if strings.HasPrefix(sLower, "i") && strings.HasSuffix(sLower, "n") && strings.Contains(sLower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
