package main

import (
	"fmt"
)

func GenDisplaceFn(a, v, d float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((1. / 2.) * a * (t * t)) + (v * t) + d
	}
}

func main() {
	var a, v, d, t float64
	fmt.Println("Please input a value for acceleration:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("error reading the acceleration: %s\n", err.Error())
	}

	fmt.Println("Please input a value for velocity:")
	_, err = fmt.Scan(&v)
	if err != nil {
		fmt.Printf("error reading the velocity: %s\n", err.Error())
	}

	fmt.Println("Please input a value for displacement:")
	_, err = fmt.Scan(&d)
	if err != nil {
		fmt.Printf("error reading the displacement: %s\n", err.Error())
	}

	displacementFn := GenDisplaceFn(a, v, d)

	fmt.Println("Please input a value for time:")
	_, err = fmt.Scan(&t)
	if err != nil {
		fmt.Printf("error reading the time: %s\n", err.Error())
	}

	fmt.Println("Displacement: ", displacementFn(t))
}
