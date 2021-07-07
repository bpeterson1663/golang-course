package main

import (
	"fmt"
	"scope/packageone"
)

var one = "One"

func main() {
	var somethingElse = "block level"
	fmt.Println(somethingElse)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("from package one: ", newString)

}

func myFunc() {
	packageone.ExportedFunc()
}
