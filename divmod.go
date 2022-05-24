package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b

	*mod = a % b
}

//This function will divide the int a and b.
//The result of this division will be stored
//in the int pointed by div.
//The remainder of this division will be stored in the int pointed by mod.
