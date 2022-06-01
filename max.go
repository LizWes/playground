package piscine

func Max(a []int) int {
	var length int = len(a)
	var maximum int = a[0]
	if length == 0 {
		maximum = 0
	} else {
		for i := 0; i < length; i++ {
			if a[i] > maximum {
				maximum = a[i]
			}
		}
	}
	return maximum
}
