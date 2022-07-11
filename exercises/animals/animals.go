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

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) PrintDesiredAttribute(attribute string) {
	switch attribute {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("That's not an existing attribute, please try again, attributes: eat, move, speak")
	}
}

func CreateAnimals() map[string]Animal {
	return map[string]Animal{
		"cow": {
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
		"bird": {
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
		"snake": {
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		},
	}
}

func ReadCommand() (string, string, error) {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("There was an error while reading your input, please try again")
	}
	command = strings.Replace(command, "\n", "", -1)
	commandComponents := strings.Split(command, " ")
	if len(commandComponents) != 2 {
		return "", "", fmt.Errorf("The format of the message is not correct, it should be 2 words separated by a string, e.g: 'cow eat'")
	}
	return commandComponents[0], commandComponents[1], nil
}

func main() {
	animals := CreateAnimals()
	for {
		animalSt, attribute, err := ReadCommand()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		animal, ok := animals[animalSt]
		if !ok {
			fmt.Println("That's not an existing animal, please try again, animals: cow, snake, bird")
			continue
		}
		animal.PrintDesiredAttribute(attribute)
	}
}
