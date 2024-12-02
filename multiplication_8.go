package main

import "fmt"

func main() {
	// Print multiplication table of 8
	fmt.Println("Multiplication table of 8:")
	for i := 1; i <= 10; i++ {
		result := 8 * i
		fmt.Printf("8 x %d = %d\n", i, result)
	}
}
