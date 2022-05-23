package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return false
	} else {
		for _, val := range s {
			if (val < 48) || (val > 57 && val < 65) || (val > 'Z' && val < 'a') || (val > 'z') {
				return false
			}
		}
	}
	return true
}
