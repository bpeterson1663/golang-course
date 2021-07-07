package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when Ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer = firstNumber*secondNumber - subtraction

	startGame(firstNumber, secondNumber, subtraction, answer)
}

func startGame(first, second, sub, answer int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", first, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", second, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", sub, prompt)
	reader.ReadString('\n')

	fmt.Println("The answer is", answer)
}
