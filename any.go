package piscine

func Any(chosenfunc func(string) bool, a []string) bool {
	for _, val := range a {
		if chosenfunc(val) == true {
			return true
		}
	}
	return false
}

// Write a function Any that returns true, for a string slice :
// if, when that string slice is passed through an f function,
// at least one element returns true.
