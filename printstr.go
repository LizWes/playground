package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, c := range s {
		if c == '\n' {
			z01.PrintRune('\\')
			z01.PrintRune('n')
		} else {
			z01.PrintRune(c)
		}
	}
}

//Write a function that prints one by one
//the characters of a string on the screen.
