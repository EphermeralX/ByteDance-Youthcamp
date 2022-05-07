package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Random number seed for timestamp not set
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
