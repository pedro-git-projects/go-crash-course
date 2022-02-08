package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press return when ready."

func main() {
	
	// seed random number
	rand.Seed(time.Now().UnixNano())
	
	// create random numbers between 2 and 10
	firstNumber := rand.Intn(8) + 2
	secondNumber := rand.Intn(8) + 2
	subtraction := rand.Intn(8) + 2
	awnser := firstNumber * secondNumber - subtraction
	
	// create a reader
	reader := bufio.NewReader(os.Stdin)
	
	// game instructions
	fmt.Println("Guess the number!")	
	fmt.Println("Choose a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// game proper
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally tought of", prompt)

	fmt.Println("Now subtract", subtraction, prompt)

	fmt.Println("The awnser is ", awnser)
}
