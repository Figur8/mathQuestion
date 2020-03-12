package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(50)
}

func answerQuestion() {
	var answer int
	fmt.Scanln(&answer)
}

func main() {
	a := createRandomNumber()
	b := createRandomNumber()
	result := a + b
	fmt.Println(a, "+", b)
	fmt.Println(result)
}
