package main

import "github.com/01-edu/z01"

type pointer struct {
	x int
	y int
}

func main() {
	points := &pointer{}

	setPoint(points)

	valueOfx := points.x
	runeOfx := []rune{}

	valueOfy := points.y
	runeOfy := []rune{}

	firstPart := "x = "
	printStr(firstPart)

	for valueOfx != 0 {
		runeOfx = append(runeOfx, rune(valueOfx%10))
		valueOfx /= 10
	}

	for i := len(runeOfx) - 1; i >= 0; i-- {
		z01.PrintRune((runeOfx[i] + 48))
	}

	thirdPart := ", y = "
	printStr(thirdPart)

	for valueOfy != 0 {
		runeOfy = append(runeOfy, rune(valueOfy%10))
		valueOfy /= 10
	}
	for i := len(runeOfy) - 1; i >= 0; i-- {
		z01.PrintRune((runeOfy[i] + 48))
	}
	z01.PrintRune('\n')
}

func setPoint(pointer *pointer) {
	pointer.x = 42
	pointer.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// Create a new directory called point.
// The code below must be copied into a
// file called main.go inside the point directory.
// The necessary changes must be applied so that the program works.
// The function setPoint() must work with int.
