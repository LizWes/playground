package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for firstDigit := '0'; firstDigit <= '9'; firstDigit++ {
		for b := firstDigit + 1; b <= '9'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(firstDigit)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if firstDigit < 55 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
	z01.PrintRune('\n')
}

//Write a function that prints, in ascending order and on a
//single line: all unique combinations of three different
//digits so that, the first digit is lower than the second,
//and the second is lower than the third.
//These combinations are separated by a comma and a space.
