package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string

	//Create slice of lenght 3
	slice := make([]int, 3)

	for input != "x" {
		fmt.Println("Enter an integer number: ")
		_, err1 := fmt.Scan(&input)
		input = strings.ToLower(input)
		number, err2 := strconv.Atoi(input)

		if err1 == nil && err2 == nil {
			slice = append(slice, number)
			fmt.Println("input: ", number)
			fmt.Println("slice: ", slice)

		}
	}

	slice = slice[3:len(slice)]
	fmt.Println("slice: ", slice)

	sort.Ints(slice)
	fmt.Println("slice: ", slice)

}
