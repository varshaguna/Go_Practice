package main

import "fmt"

func main() {
	// Initialize sum variable
	sum := 0

	// Calculate the sum of numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		sum += i
	}

	// Print the sum
	fmt.Println("Sum of numbers from 1 to 10 is:", sum)
}
