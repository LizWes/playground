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

//  Write a function IsSorted that returns true,
// if the slice of int is sorted, otherwise returns false.
// The function passed in the parameter returns a positive
// int if a (the first argument) is greater than to b (the second argument),
// it returns 0 if they are equal and it returns a negative int otherwise.
// To do your testing you have to write your own f function.
