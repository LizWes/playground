package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return RecursivePower(nb, power-1) * nb
}

// Write a recursive function that returns the value
// of nb to the power of power.
// Negative powers will return 0. Overflows do not have
// to be dealt with.
// for is forbidden for this exercise.
