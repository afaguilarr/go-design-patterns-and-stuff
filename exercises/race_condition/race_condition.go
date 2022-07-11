package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int
	firstSetOfNumbers := []int{5, 10, 15}
	secondSetOfNumbers := []int{6, 12, 18}

	// This first goroutine should modify the sum variable to be 30
	go func() {
		for _, num := range firstSetOfNumbers {
			// This will generate some discrepancies depending on the program interleavings
			time.Sleep(time.Millisecond * 12)
			sum += num
		}
	}()

	// This goroutine should modify the sum variable to be 36, to a total of 66
	go func() {
		for _, num := range secondSetOfNumbers {
			// This will generate some discrepancies depending on the program interleavings
			time.Sleep(time.Millisecond * 12)
			sum += num
		}
	}()
	// The sleeping time of the 3 goroutines should end at the same time but that isn't the case,
	// sometimes the main goroutine will end first, sometimes both goroutines above will end first,
	// and sometimes only one of the goroutines above will reach the end before the main goroutine.
	time.Sleep(time.Millisecond * 36)
	// That's why executing this code 10 times will probably show different results such as 33, 48, 51, 66
	fmt.Println("Final result: ", sum)
}
