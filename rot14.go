package piscine

func Rot14(s string) string {
	var output string
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			if letter > 'l' {
				output += string(letter - 12)
			} else {
				output += string(letter + 14)
			}
		} else if letter >= 'A' && letter <= 'Z' {
			if letter > 'L' {
				output += string(letter - 12)
			} else {
				output += string(letter + 14)
			}
		} else {
			output += string(letter)
		}
	}
	return (output)
}
