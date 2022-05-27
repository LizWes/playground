package piscine

func IsSorted(chosenfunc func(a1, a2 int) int, tab []int) bool {
	var sorted bool = true
	for i := 0; i < len(tab)-1; i++ {
		if chosenfunc(tab[i], tab[i+1]) > 0 {
			sorted = false
		}
	}
	if sorted == false {
		sorted = true
		for i := 0; i < len(tab)-1; i++ {
			if chosenfunc(tab[i], tab[i+1]) < 0 {
				sorted = false
			}
		}
	}
	return sorted
}
