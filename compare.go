package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	} else {
		return 1
	}
}

//Write a function that behaves like the Compare function.
