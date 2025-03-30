package helpers

import (
	"fmt"
)

var reset string = "\033[0m"

func PrintGreen(text string) {
	green := "\033[32m"
	fmt.Println(green + text + reset)
}

func PrintRed(text string) {
	red := "\033[31m"
	fmt.Println(red + text + reset)
}
