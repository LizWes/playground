package piscine

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

// Write a function that takes two pointers to
// an int (*int) and swaps their contents.
