package piscine

func ForEach(f func(int), a []int) {

	for _, val := range a {
		f(val)
	}
}

// Write a function ForEach that, for an int slice,
// applies a function on each element of that slice.
