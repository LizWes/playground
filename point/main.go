package main

import "github.com/01-edu/z01"

type pointer struct {
	x int
	y int
}

func setPoint(pointer *pointer) {
	pointer.x = 42
	pointer.y = 21
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

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
