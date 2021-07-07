package packageone

import (
	"fmt"
)

var packageVar = "Exported from packageone"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, packageVar)
}
