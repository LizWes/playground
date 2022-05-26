package piscine

func TrimAtoi(s string) int {
	i := 0

	for _, let := range s {
		if let >= '1' && let <= '9' {
			if let == '0' {
				i += 0
			} else if let == '1' {
				i += 1
			} else if let == '2' {
				i += 2
			} else if let == '3' {
				i += 3
			} else if let == '4' {
				i += 4
			} else if let == '5' {
				i += 5
			} else if let == '6' {
				i += 6
			} else if let == '7' {
				i += 7
			} else if let == '8' {
				i += 8
			} else if let == '9' {
				i += '9'
			}
			i *= 10
		}
	}
	num := i / 10
	mind := 100
	lin := 0
	for index, let := range s {
		if let == '-' {
			mind = index
		}
		if let > '0' && let <= '9' {
			lin = index
			if index > mind {
				return -num
			} else if lin > 0 {
				return num
			}
		}

	}
	return num
}

// Write a function which prints the digits of an int
// passed in parameter in ascending order. All possible
// values of type int have to go through, excluding
// negative numbers. Conversion to int64 is not allowed.
