package main

import "fmt"

func main() {
	// Declare an array
	var arr [5]int

	// Assign values to an array
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// Print array values
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
