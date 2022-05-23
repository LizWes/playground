package piscine

func IterativeFactorial(nb int) int {
	max := 21
	if nb < 0 {
		return 0
	} else {
		result := 1

		for i := 1; i <= nb; i++ {
			if i > max {
				return 0
			}
			result = result * i
		}
		return result
	}
}