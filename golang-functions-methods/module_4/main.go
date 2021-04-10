package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("Food eaten: grass")
}

func (c Cow) Move() {
	fmt.Println("Locomotion method: walk")
}

func (c Cow) Speak() {
	fmt.Println("Spoken sound: moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("Food eaten: worms")
}

func (b Bird) Move() {
	fmt.Println("Locomotion method: fly")
}

func (b Bird) Speak() {
	fmt.Println("Spoken sound: peep")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("Food eaten: mice")
}

func (s Snake) Move() {
	fmt.Println("Locomotion method: slither")
}

func (b Snake) Speak() {
	fmt.Println("Spoken sound: hsss")
}

var animalsMap = make(map[string]Animal)

func main() {
	fmt.Println("Enter in one line name of command ['newanimal', 'query'], name of animal and either type of animal to create a new animal: ['cow', 'bird', 'snake'] or name of the information requested about the animal ['eat', 'move', 'speak'] separated by whitespace:")
	for {
		var command, animalName, request string
		fmt.Printf("> ")
		fmt.Scan(&command, &animalName, &request)
		switch command {
		case "newanimal":
			createAnimal(animalName, request)
		case "query":
			queryAnimal(animalName, request)
		default:
			fmt.Println("Unexpected command provided, first argument of the string must be 'newanimal' or 'query'.")
		}

	}
}

func createAnimal(name, animalType string) {
	var animal Animal
	switch animalType {
	case "cow":
		animal = Cow{name: name}
	case "bird":
		animal = Bird{name: name}
	case "snake":
		animal = Snake{name: name}
	default:
		fmt.Println("Incorrect request: requested animal can't be created, available animals: [cow, bird, snake]")
		return
	}
	animalsMap[name] = animal
	fmt.Println("Created it!")
}

func queryAnimal(name, infoRequest string) {
	animal, ok := animalsMap[name]
	if !ok {
		fmt.Printf("No animal with name %s found\n", name)
		return
	}
	switch infoRequest {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Unknown reques, supported requests: ['eat', 'move', 'speak']")
	}

}
