package utils

import (
	"math/rand"
	"time"
)

// Generate random four-digit number
func GenerateRandom4DigitNumber() int {
	rand.Seed(time.Now().UnixNano())

	low := 1000
	high := 9999
	return low + rand.Intn(high-low)
}
