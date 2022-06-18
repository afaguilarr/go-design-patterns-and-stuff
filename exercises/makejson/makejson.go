package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string
	person := map[string]string{}

	fmt.Println("Please input your name:")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Printf("Something went wrong when reading your name, please try again. Error: %v\n", err)
		return
	}

	fmt.Println("Please input your address:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		address = scanner.Text()
	}
	if address == "" {
		fmt.Printf("Something went wrong when reading your address, please try again.")
		return
	}

	person["name"] = name
	person["address"] = address

	jsonByteArray, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Something went wrong when marshaling the JSON, please try again. Error: %v\n", err)
		return
	}
	fmt.Printf("The resulting JSON is: %s\n", string(jsonByteArray))
}
