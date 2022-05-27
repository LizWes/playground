package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, val := range a {
		result[i] = f(val)
	}
	return result
}

// Write a function Map that, for an int slice,
// applies a function of this type func(int) bool
//on each element of that slice and returns a slice of all the return values.
