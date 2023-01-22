package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	slice := generateSlice(30)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	slice = bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}
