package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main()  {
	var chosenAnimal Animal
	for {
		fmt.Print("Please enter request your request in the form of 'name,action' \n")
		fmt.Print("Example: cow,eat \n")
		fmt.Print(">")

		var userInput string
		invalidInput := false

		_, _ = fmt.Scan(&userInput)

		splitUserInput := strings.Split(userInput, ",")

		animal := splitUserInput[0]
		action := splitUserInput[1]

		switch animal {
		case "cow":
			chosenAnimal = Animal{"grass", "walk", "moo"}
		case "bird":
			chosenAnimal = Animal{"worms", "fly", "peep"}
		case "snake":
			chosenAnimal = Animal{"mice", "slither", "hsss"}
		default:
			invalidInput = true
			fmt.Println("Please provide a correct animal name")
		}

		if !invalidInput {
			switch action {
			case "eat":
				chosenAnimal.Eat()
			case "move":
				chosenAnimal.Move()
			case "speak":
				chosenAnimal.Speak()
			default:
				fmt.Println("Please provide a correct animal action")
			}
		}
	}
}
