package main

import (
	"fmt"
	"sort"
)

// Function to find the median
func findMedian(arr []int) float64 {
	// Sort the array
	sort.Ints(arr)

	n := len(arr)

	// If the number of elements is odd, return the middle element
	if n%2 != 0 {
		return float64(arr[n/2])
	}

	// If the number of elements is even, return the average of the two middle elements
	mid1 := arr[(n/2)-1]
	mid2 := arr[n/2]
	return float64(mid1+mid2) / 2.0
}

func main() {
	var n int

	// Ask the user for the size of the array
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("The array size must be greater than zero.")
		return
	}

	// Create an array of size n
	arr := make([]int, n)

	// Get array elements from the user
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Find and print the median
	median := findMedian(arr)
	fmt.Printf("The median is: %.2f\n", median)
}
