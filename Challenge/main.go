package main

import (
	"challenge/packageone"
	"fmt"
)

var myVar = "package level variable"

func main() {
	var blockVar = "block level for main"
	fmt.Println()
	packageone.PrintMe(myVar, blockVar)
}
