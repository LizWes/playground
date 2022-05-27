package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	} else {
		tab := make([]int, max-min)
		j := 0

		for i := min; i < max; i++ {
			tab[j] = i
			j++
		}
		return tab
	}
}

// Write a function that takes an int min
// and an int max as parameters. The function
// must return a slice of ints with all the values between min and max.
// Min is included, and max is excluded.
// If min is greater than or equal to max, a nil slice is returned.
// append is not allowed for this exercise.
