package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myString = new(string)

	fmt.Println("Please input a floating number, using point as a separator:")
	_, err := fmt.Scan(myString)
	if err != nil {
		fmt.Println("There was an error in the input, please try again")
		return
	}

	myFloat, err := strconv.ParseFloat(*myString, 64)
	if err != nil {
		fmt.Println("There was an error while trying to parse the float number, please try again")
		return
	}

	myInt := int(myFloat)
	fmt.Printf("The truncated version of the number you used as an input is: %v\n", myInt)
}
