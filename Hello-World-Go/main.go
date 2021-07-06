package main

import (
	"fmt"
	"myApp/doctor"
)

func main() {
	var whatToSay string
	whatToSay = doctor.Intro()
	// whatToSay := "Hellow world as an argument" //shorthand to infer type
	// sayHelloWorld(whatToSay)
	fmt.Println(whatToSay)
}

// func sayHelloWorld(whatToSay string) {
// 	fmt.Println(whatToSay)
// }
