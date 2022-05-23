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
