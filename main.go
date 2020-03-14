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

func answerQuestion() int {
	var answer int
	fmt.Println("What's your answer?")
	fmt.Scanln(&answer)
	return answer
}

func main() {
	a := createRandomNumber()
	b := createRandomNumber()
	fmt.Println("How much is: ", a, "+", b)
	result := a + b
	if result == answerQuestion() {
		fmt.Println("You are right, the result is: ", result)
	} else {
		fmt.Println("You are wrong, the result is: ", result)
	}
}
