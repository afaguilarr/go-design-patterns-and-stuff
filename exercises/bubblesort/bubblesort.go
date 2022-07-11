package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	firstElement, secondElement *int
}

func SwapFn() func(p Pair) int {
	numberOfSwaps := 0

	Swap := func(p Pair) int {
		if *(p.firstElement) > *(p.secondElement) {
			temp := *(p.firstElement)
			*(p.firstElement) = *(p.secondElement)
			*(p.secondElement) = temp
			numberOfSwaps++
		}
		return numberOfSwaps
	}

	return Swap
}

func main() {
	line, err := readLine()
	if err != nil {
		fmt.Printf("There was an error while reading the elements: %s\n", err.Error())
		return
	}

	elements, err := parseLine(line)
	if err != nil {
		fmt.Printf("There was an error while parsing the elements: %s\n", err.Error())
		return
	}

	BubbleSort(elements)
	fmt.Printf("The sorted list is: %v\n", elements)
}

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input up to 10 integer numbers separated by single-space characters:")
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("There was an error while reading your input: %s", err.Error())
	}

	line = strings.Replace(line, "\n", "", -1)
	return line, nil
}

func parseLine(line string) ([]int, error) {
	strElements := strings.Split(line, " ")
	if len(strElements) > 10 {
		return nil, fmt.Errorf("The list can only have 10 elements at most")
	}

	elements := make([]int, len(strElements))
	for i, strEl := range strElements {
		el, err := strconv.ParseInt(strEl, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("There was an error parsing the numbers in the list: %s", err.Error())
		}
		elements[i] = int(el)
	}

	return elements, nil
}

func BubbleSort(slice []int) {
	for {
		var numberOfSwaps int
		swap := SwapFn()
		for i := range slice {
			if i == 0 {
				continue
			}

			p := Pair{firstElement: &slice[i-1], secondElement: &slice[i]}
			numberOfSwaps = swap(p)
		}
		if numberOfSwaps == 0 {
			break
		}
	}
}
