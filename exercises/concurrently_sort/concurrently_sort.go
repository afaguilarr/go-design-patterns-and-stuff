package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input only integer numbers separated by single-space characters:")
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("There was an error while reading your input: %s", err.Error())
	}

	line = strings.Replace(line, "\n", "", -1)
	return line, nil
}

func parseLine(line string) ([]int, error) {
	strElements := strings.Split(line, " ")
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

func concurrentlySort(elements []int) []int {
	lists := getPartitionedLists(elements)
	var wg sync.WaitGroup
	for _, l := range lists {
		wg.Add(1)
		go sortList(l, &wg)
	}
	wg.Wait()
	newElements := mergeLists(lists)
	return newElements
}

func mergeLists(lists [][]int) []int {
	if len(lists) != 2 {
		firstList := mergeLists(lists[0:2])
		secondList := mergeLists(lists[2:4])
		return mergeLists([][]int{firstList, secondList})
	}

	list := []int{}
	for {
		if len(lists[0]) == 0 {
			for _, e := range lists[1] {
				list = append(list, e)
			}
			break
		}
		if len(lists[1]) == 0 {
			for _, e := range lists[0] {
				list = append(list, e)
			}
			break
		}
		min := lists[0][0]
		minIndex := 0
		if lists[1][0] < lists[0][0] {
			min = lists[1][0]
			minIndex = 1
		}
		list = append(list, min)
		lists[minIndex] = lists[minIndex][1:]
	}
	return list
}

func getPartitionedLists(elements []int) [][]int {
	var extraElements []int
	mod := len(elements) % 4
	if mod != 0 {
		extraElements = elements[len(elements)-mod : len(elements)]
	}

	lists := make([][]int, 4)
	p := int(len(elements) / 4)
	for i := 0; i < 4; i++ {
		lists[i] = make([]int, p)
		copy(lists[i], elements[i*p:(i+1)*p])
	}

	if len(extraElements) != 0 {
		for i, e := range extraElements {
			lists[i] = append(lists[i], e)
		}
	}
	return lists
}

func sortList(list []int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	fmt.Println("Sub-array before sorting: ", list)
	sort.Ints(list)
	fmt.Println("Sub-array after sorting: ", list)
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

	elements = concurrentlySort(elements)
	fmt.Printf("The sorted list is: %v\n", elements)
}
