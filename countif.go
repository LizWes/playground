package piscine

func CountIf(chosenfunc func(string) bool, tab []string) int {
	count := 0
	for _, val := range tab {
		if chosenfunc(val) == true {
			count++
		}
	}
	return count
}
