package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float32
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("What is your name?")
	prompt()

	user.Age = readInt("How old are you")

	user.FavoriteNumber = readFloats("What is your favorite number")
	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.2f\n", user.UserName, user.Age, user.FavoriteNumber)
}

func prompt() {
	fmt.Print("=> ")

}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a valid name")
		} else {
			return userInput

		}
	}

}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please Enter a valid number")
		} else {
			return num

		}
	}
}

func readFloats(s string) float32 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 32)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return float32(num)

		}
	}
}
