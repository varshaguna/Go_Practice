package main

import "fmt"

func main() {
	// Loop through numbers from 1 to 100
	for i := 1; i <= 100; i++ {
		// Perform division and print the result
		fmt.Printf("10 / %d = %.2f\n", i, 10.0/float64(i))
	}
}
