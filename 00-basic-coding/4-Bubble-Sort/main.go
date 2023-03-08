package main

import (
	"fmt"
	"time"
)

// []int{5, 2, 1, 8, 4}
// i j
// 0 0 {2, 5, 1, 8, 4}
// 0 1 {2, 1, 5, 8, 4}
// 0 3 {2, 1, 5, 4, 8}

// 1 0 {1, 2, 5, 4, 8}
// 1 2 {1, 2, 4, 5, 8}

func BubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSort2(arr []rune) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {

	start := time.Now()

	// Example usage
	// arr := []int{5, 2, 1, 8, 4}
	// arr := []int{5, 2, 1, 8, 4, 6, 9, 20, 12, 19, 89, 100, 201, 200}
	// arr := []int{98, 99, 97}
	arr := []rune{98, 99, 97}
	BubbleSort2(arr)
	fmt.Println(arr) // Output: [1 2 4 5 8]

	elapsed := time.Now().Sub(start)
	fmt.Printf("Function took %s", elapsed)
}
