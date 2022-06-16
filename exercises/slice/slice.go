package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sl := make([]int, 3)

	counter := 0
	for {
		input := new(string)
		fmt.Println("Please input a number, if you want to close the program, please input the character 'X':")
		_, err := fmt.Scan(input)
		if err != nil {
			fmt.Printf("There was an error while trying to read the input: %v\n", err.Error())
			continue
		}

		inputInt, err := strconv.ParseInt(*input, 10, 64)
		if err != nil {
			lowerInput := strings.ToLower(*input)
			if lowerInput == "x" {
				fmt.Println("Bye bye")
				return
			} else {
				fmt.Println("The input is unknown, please try again")
				continue
			}
		}

		alterSlice(&sl, &counter, int(inputInt))
		fmt.Printf("Resulting slice: %v\n", sl)
	}
}

func alterSlice(sl *[]int, counter *int, inputInt int) {
	if *counter < 3 {
		for index, el := range *sl {
			if el == 0 {
				(*sl)[index] = inputInt
				break
			}
		}
		*counter++
	} else {
		*sl = append(*sl, inputInt)
	}
	sort.Ints(*sl)
}
