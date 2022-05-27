package piscine

func Any(chosenfunc func(string) bool, a []string) bool {
	for _, val := range a {
		if chosenfunc(val) == true {
			return true
		}
	}
	return false
}
