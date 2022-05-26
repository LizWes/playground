package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	order := [10]int{}
	if n >= 0 && n <= 9 {
		z01.PrintRune(rune(n + 48))
	} else {
		for {
			if n == 0 {
				break
			}
			order[n%10]++
			n = n / 10
		}
		for index, num := range order {
			if num != 0 {
				for i := 0; i < num; i++ {
					z01.PrintRune(rune(index + 48))
				}
			}
		}
	}
}

// Write a function which prints the digits of an int
// passed in parameter in ascending order. All possible
// values of type int have to go through, excluding negative numbers.
// Conversion to int64 is not allowed.
