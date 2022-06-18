package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fname, lname string
}

func (n *Name) SetAttributes(line string, lineNumber int) error {
	namesSlice := strings.Split(line, " ")
	if len(namesSlice) != 2 {
		return fmt.Errorf("The format of the line with number '%v' and content '%s' is not expected", lineNumber, line)
	}
	n.fname = namesSlice[0]
	n.lname = namesSlice[1]
	return nil
}

func main() {
	var fileName string

	fmt.Println("Please input the name of the file:")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Printf("Something went wrong when reading the file name, please try again. Error: %v\n", err)
		return
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Something went wrong when reading the file, please try again. Error: %v\n", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	names := make([]Name, len(lines))
	for index, line := range lines {
		err = names[index].SetAttributes(line, index+1)
		if err != nil {
			fmt.Printf("There was an error processing the lines: %s\n", err.Error())
			return
		}
	}

	for _, name := range names {
		fmt.Printf("First name: %s\n", name.fname)
		fmt.Printf("Last name: %s\n", name.lname)
	}
}
