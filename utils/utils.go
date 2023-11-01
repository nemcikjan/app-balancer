package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	fmt.Println("Seeding ...")
	rand.Seed(time.Now().UnixNano())
}

func RandomNumber(max int) int {

	// Generate a random integer between 0 and 100
	return rand.Intn(max + 1)
}

func RandomString(length int) string {
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
