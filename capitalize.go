package piscine

func Capitalize(s string) string {
	var result string
	isFirstCapital := 0

	for _, val := range s {
		if (val >= 'a') && (val <= 'z') {
			if isFirstCapital == 0 {
				result += string(val - 32)
			} else {
				result += string(val)
			}
			isFirstCapital = 1

		} else if (val >= 'A') && (val <= 'Z') {
			if isFirstCapital == 0 {
				result += string(val)
			} else {
				result += string(val + 32)
			}
			isFirstCapital = 1

		} else if (val >= '0') && (val <= '9') {
			result += string(val)
			isFirstCapital = 1
		} else {
			result += string(val)
			isFirstCapital = 0
		}
	}
	return result
}
