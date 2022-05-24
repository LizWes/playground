package piscine

func RecursiveFactorial(nb int) int {
	max := 21
	if nb < 0 || nb > max {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb >= 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}

// Write a recursive function that returns the
// factorial of the int passed as parameter.
// Errors (non possible values or overflows) will return 0.
// for is forbidden for this exercise.
