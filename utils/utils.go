package utils

import (
	"log"
	"math/rand"
	"time"
)

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("\n%v: %v\n", msg, time.Since(start))
}

func GenerateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size) - rand.Intn(size)
	}
	return slice
}
