package piscine

import "github.com/01-edu/z01"

func IsNegative(number int) {
	if (number) >= 0 {
		z01.PrintRune('F')
		z01.PrintRune('\n')

	} else {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	}
}

// Write a function that prints 'T' (true) on
// a single line if the int passed as parameter
// is negative, otherwise it prints 'F' (false).
