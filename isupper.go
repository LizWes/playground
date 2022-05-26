package piscine

func IsUpper(s string) bool {
	for _, val := range s {
		if val < 'A' || val > 'Z' {
			return false
		}
	}
	return true
}

// Write a function that returns true if the string
// passed as parameter contains only uppercase characters, otherwise, returns false.
