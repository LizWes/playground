package piscine

func BasicAtoi(s string) int {
	var count, sum int
	for _, value := range s {
		if value >= '0' && value <= '9' {
			count = 0
			for i := '0'; 1 < value; i++ {
				count++
			}
			sum = sum*10 + count
		}
	}
	return sum
}
