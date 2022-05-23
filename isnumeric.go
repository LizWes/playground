package piscine

func IsNumeric(s string) bool {
	for _, val := range s {
		if val < 48 || val > 57 {
			return false
		}
	}
	return true
}
