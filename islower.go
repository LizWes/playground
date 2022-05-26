package piscine

func IsLower(s string) bool {
	for _, val := range s {
		if val < 'a' || val > 'z' {
			return false
		}
	}
	return true
}

// Write a function that returns true if the string passed
// as the parameter contains only lowercase characters, otherwise, returns false.
