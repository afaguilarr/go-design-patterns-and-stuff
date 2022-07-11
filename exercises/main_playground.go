package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Swap(s []int, i int) {
	temp := s[i]
	s[i] = s[i+1]
	s[i+1] = temp
}

func BubbleSort(s []int, wg *sync.WaitGroup) {
	fmt.Printf("Go routine sorting slice : %v\n", s)
	size := len(s)

	for countdown := size; countdown >= 1; countdown-- {
		for i := 0; i < countdown-1; i++ {
			if s[i] > s[i+1] {
				Swap(s, i)
			}
		}
	}
	fmt.Printf("sorted slice : %v\n", s)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// input
	fmt.Print("Enter a serie of numbers : ")
	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')

	numbers := []int{}
	for _, f := range strings.Fields(line) {
		myNumber, _ := strconv.Atoi(f)
		numbers = append(numbers, myNumber)
	}

	fmt.Println(numbers)

	var size = len(numbers)
	var nParts = 4
	var partSize = int(math.Ceil(float64(size) / 4))
	var parts [][]int

	fmt.Printf("numbers : %d, parts : %d, partSize : %d\n", size, nParts, partSize)

	// divide slice into nParts slices
	for i := 0; i < size; i += partSize {
		end := i + partSize

		if end > size {
			end = size
		}

		parts = append(parts, numbers[i:end])
	}
	//fmt.Printf("%v\n", parts)

	//  sort all parts in go routines
	for i := 0; i < len(parts); i++ {
		wg.Add(1)
		go BubbleSort(parts[i], &wg)
	}

	wg.Wait()
	//fmt.Printf("%v\n", parts)

	// merge all parts into numbersSorted slice
	numbersSorted := []int{}
	for i := 0; i < len(parts); i++ {
		s1 := parts[i]

		for _, x := range s1 {
			//fmt.Println(x)
			ind := sort.Search(len(numbersSorted), func(i int) bool { return numbersSorted[i] > x })
			fmt.Println("wtf1: ", numbersSorted)
			numbersSorted = append(numbersSorted, 0)
			fmt.Println("wtf2: ", numbersSorted)
			copy(numbersSorted[ind+1:], numbersSorted[ind:])
			fmt.Println("wtf3: ", numbersSorted)
			numbersSorted[ind] = x
			fmt.Println("wtf4: ", numbersSorted)
		}
		//fmt.Println(numbersSorted)

	}

	fmt.Println(numbersSorted)

	//BubbleSort(numbers)
	//fmt.Println(numbers)
}
