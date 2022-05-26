package piscine

func NRune(str string, num int) rune {
	runes := []rune(str)
	len := 0
	for i := range runes {
		len = i
	}
	if num-1 >= 0 && num-1 <= len {
		return runes[num-1]
	}
	return 0
}

// Write a function that returns the nth rune
// of a string. If not possible, it returns 0.
