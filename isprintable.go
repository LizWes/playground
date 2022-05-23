package piscine

func IsPrintable(s string) bool {
	for _, val := range s {
		if val < 21 {
			return false
		}
	}
	return true
}
