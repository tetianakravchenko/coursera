package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

var animals = map[string]Animal{
	"cow": Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	},
	"bird": Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	},
	"snake": Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	},
}

func main() {
	fmt.Println("Enter in one line name of an enimal (either 'cow', 'bird', or 'snake') and name of the information requested about the animal (either 'eat', 'move', or 'speak') separated by whitespace:")
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		inputTrim := strings.TrimSuffix(input, "\n")
		inputArr := strings.Split(inputTrim, " ")

		if len(inputArr) != 2 {
			fmt.Println("Not enough data provided. ")
		}
		_, ok := animals[inputArr[0]]
		if !ok {
			fmt.Println("Incorrect request: requested animal doesn't exist, available animals: []")
		}
		var inputAnimal Animal = animals[inputArr[0]]
		switch inputArr[1] {
		case "eat":
			inputAnimal.Eat()
		case "move":
			inputAnimal.Move()
		case "speak":
			inputAnimal.Speak()
		default:
			fmt.Println("Unexpected request about the animal, available information: [eat, move, speak]")
		}

	}
}

func (a Animal) Eat() {
	fmt.Printf("Food eaten: %s\n", a.food)
}

func (a Animal) Move() {
	fmt.Printf("Locomotion method: %s\n", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("Spoken sound: %s\n", a.locomotion)
}
