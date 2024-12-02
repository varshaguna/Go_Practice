package main

import "fmt"

func main() {
	var x, y int

	// Take input for x and y from the user
	fmt.Print("Enter the starting number (x): ")
	fmt.Scan(&x)
	fmt.Print("Enter the ending number (y): ")
	fmt.Scan(&y)

	// Check if x is less than or equal to y
	if x <= y {
		fmt.Println("Numbers from x to y:")
		for i := x; i <= y; i++ {
			fmt.Println(i)
		}
	} else {
		// If x > y, count down from x to y
		fmt.Println("Numbers from x to y in reverse order:")
		for i := x; i >= y; i-- {
			fmt.Println(i)
		}
	}
}
