package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}

//This function will divide the int a and b.
//The result of this division will be stored
//in the int pointed by a.
//The remainder of this division will be stored in the int pointed by b.
