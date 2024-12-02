package main

import "fmt"

func main() {
	fmt.Println("Numbers from 1 to 10, except 5:")

	for i := 1; i <= 10; i++ {
		// Skip the number 5
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
