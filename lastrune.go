package piscine

func LastRune(s string) rune {
	converter := []rune(s)
	return rune(converter[len(s)-1])
}
