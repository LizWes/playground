package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNum(num int) {
	i := '0'
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= num%10; j++ {
		i++
	}
	fmt.Println("This is i in first loop")
	fmt.Println(i)

	for j := -1; j >= num%10; j-- {
		i++
	}

	fmt.Println("This is i in second loop %v", i)
	fmt.Println("This is the result")

	if num/10 != 0 {
		PrintNum(num / 10)
	} else {
		z01.PrintRune(i)
	}

}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}
