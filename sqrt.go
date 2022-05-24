package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := nb / 2; i >= 1; i-- {
		tmp := i * i
		if tmp == nb {
			return i
		}
	}
	return 0
}

// Write a function that returns the square root
// of the int passed as parameter, if that square
// root is a whole number. Otherwise it returns 0.
