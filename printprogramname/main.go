package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	stringA := []rune(os.Args[0])
	size := len(stringA)

	for i := 0; i < size; i++ {
		if stringA[i] != '.' && stringA[i] != '/' {
			z01.PrintRune(stringA[i])
		}
	}
	z01.PrintRune('\n')
}
