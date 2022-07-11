package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow string
type Bird string
type Snake string

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func ReadCommand() (string, string, string, error) {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		return "", "", "", fmt.Errorf("There was an error while reading your input, please try again")
	}
	command = strings.Replace(command, "\n", "", -1)
	commandComponents := strings.Split(command, " ")
	if len(commandComponents) != 3 {
		return "", "", "", fmt.Errorf("The format of the message is not correct, it should be 3 words separated by a string, e.g: 'query mary eat'")
	}
	return commandComponents[0], commandComponents[1], commandComponents[2], nil
}

func InsertAnimal(name, attribute string, animals map[string]Animal) error {
	switch attribute {
	case "cow":
		animals[name] = Cow(name)
	case "bird":
		animals[name] = Bird(name)
	case "snake":
		animals[name] = Snake(name)
	default:
		return fmt.Errorf("The animal type is not valid, valid animals: cow, bird, snake")
	}
	fmt.Println("Created it!")
	return nil
}

func QueryAnimalAttribute(name, attribute string, animals map[string]Animal) error {
	a, ok := animals[name]
	if !ok {
		return fmt.Errorf("Animal not found, please input an existent name")
	}
	err := PrintDesiredAttribute(a, attribute)
	if err != nil {
		return err
	}
	return nil
}

func PrintDesiredAttribute(a Animal, attribute string) error {
	switch attribute {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		return fmt.Errorf("That's not an existing attribute, please try again, attributes: eat, move, speak")
	}
	return nil
}

func main() {
	animals := map[string]Animal{}
	for {
		command, name, attribute, err := ReadCommand()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		switch command {
		case "newanimal":
			err = InsertAnimal(name, attribute, animals)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "query":
			err = QueryAnimalAttribute(name, attribute, animals)
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}
