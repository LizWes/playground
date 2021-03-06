package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, val := range s {
		if (val >= 65 && val <= 90) || (val >= 97 && val <= 122) {
			counter++
		}
	}
	return counter
}

// Write a function that counts the letters of a string
// and returns the count.
// The letters are from the Latin alphabet list only,
// any other characters, symbols or empty spaces shall not be counted.
