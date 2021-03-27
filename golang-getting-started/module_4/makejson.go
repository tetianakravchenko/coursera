package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	person := make(map[string]string, 1)

	var name string
	fmt.Println("Enter a name:")
	fmt.Scan(&name)
	person["name"] = name

	fmt.Println("Enter an address:")
	reader := bufio.NewReader(os.Stdin)
	adressInput, _ := reader.ReadString('\n')
	address := strings.TrimSuffix(adressInput, "\n")
	person["address"] = address

	barr, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(barr))
}
