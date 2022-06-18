package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname, lname string
}

func buildNameList(fileName string) []Name {
	nameList := make([]Name, 0, 20)
	f, err := os.Open(fileName)
	bArr := make([]byte, 42)

	nb, readErr := f.Read(bArr)
	if err == nil && readErr == nil {
		for nb != 0 {
			fullName := strings.Split(string(bArr), " ")
			first, last := fullName[0], fullName[1]
			nameList = append(nameList, Name{first, last})
			nb, readErr = f.Read(bArr)
		}
	} else {
		fmt.Println("Oh no! We could not read from this file. Please try again.")
	}

	return nameList
}

func main() {
	var fileName string

	fmt.Print("What is the name of your file? ")
	fmt.Scan(&fileName)

	nameList := buildNameList(fileName)
	fmt.Println(nameList)
	for _, name := range nameList {
		fmt.Printf("First name: %s, Last name: %s", name.fname, name.lname)
	}
}
