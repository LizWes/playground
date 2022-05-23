package piscine

func Lenght(str string) int {
	counter := 0
	for range str {
		counter++
	}
	return counter
}

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	ss := []rune(s)
	bb := []rune(toFind)
	lenght := Lenght(toFind)
	ind := -1
	counter := 0
	for i := 0; i < Lenght(s); i++ {
		if ss[i] == bb[counter] {
			if counter == 0 {
				ind = i
			}
			counter++
			if counter == lenght {
				return ind
			}
		} else {
			if counter > 0 {
				i--
			}
			counter = 0
			ind = -1
		}
	}
	return ind
}
