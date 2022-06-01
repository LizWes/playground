package piscine

func Unmatch(a []int) int {
	var length int = len(a)
	var result int = -1
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	if length == 1 {
		return result
	} else {
		for i := 0; i < length; i += 2 {
			if i+1 >= length {
				return a[i]
			} else {
				if a[i] != a[i+1] {
					return a[i]
				}
			}
		}
	}
	return result
}
