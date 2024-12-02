package main

import "fmt"

func sumOfFactors(n int) int {
	sum := 0
	// Loop to find all factors of n
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i // Add factor to sum
		}
	}
	return sum
}

func main() {
	var num int

	// Ask the user to input a number
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Calculate and print the sum of factors
	result := sumOfFactors(num)
	fmt.Printf("Sum of factors of %d is: %d\n", num, result)
}
