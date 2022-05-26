package piscine

func ToUpper(s string) string {
	var result string

	for _, val := range s {
		if val >= ('a') && val <= 'z' {
			result += string(val - 32)
		} else {
			result += string(val)
		}
	}
	return result
}

//Write a function that capitalizes each letter in a string.
