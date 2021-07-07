package packageone

import (
	"fmt"
)

var privateVar = "I am private"

var PublicVar = "I am public or exported" // Capital letter exports variable and makes it public

func notExported() {

}

func ExportedFunc() {
	fmt.Println("I am exported")
}
