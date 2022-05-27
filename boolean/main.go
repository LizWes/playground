package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args
	lengthOfArg := len(input)

	if isEven(lengthOfArg) == 1 {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) int {
	if (nbr % 2) == 1 {
		return 1
	} else {
		return 0
	}
}

// Create a new directory called boolean.
// The code below must be copied into a file called main.go inside of the boolean directory.
// The necessary changes must be applied for the program to work
