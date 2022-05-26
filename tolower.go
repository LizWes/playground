package piscine

func ToLower(s string) string {
	var result string

	for _, val := range s {
		if val >= ('A') && val <= 'Z' {
			result += string(val + 32)
		} else {
			result += string(val)
		}
	}
	return result
}
