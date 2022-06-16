package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Please input a string:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("There was an issue while reading the input, please try again")
		return
	}

	input = strings.ToLower(input)

	// These 2 if statements could have been put together in an if statement with multiple conditions
	if input[0] != 'i' { // Same as using HasPrefix
		fmt.Println("Not Found!")
		return
	}

	if input[len(input)-1] != 'n' { // Same as using HasSuffix
		fmt.Println("Not Found!")
		return
	}

	for i := 1; i < len(input)-1; i++ { // Contains could have been used to check this too
		if input[i] == 'a' {
			fmt.Println("Found!")
			return
		}
	}

	fmt.Println("Not Found!")
}
