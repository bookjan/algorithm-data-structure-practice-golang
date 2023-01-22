package main

import (
	"fmt"
	"time"

	A "practice/algorithms"
	U "practice/utils"
)

func main() {
	unSorted := U.GenerateSlice(2000)
	fmt.Println("\n--- Unsorted --- \n\n", unSorted)

	var start time.Time

	start = time.Now()
	A.BubbleSort(unSorted)
	fmt.Printf("\n%v: %v", "BubbleSort", time.Since(start))

	start = time.Now()
	A.QuickSort(unSorted, 0, len(unSorted)-1)
	fmt.Printf("\n%v: %v", "QuickSort", time.Since(start))

	fmt.Println()
}
