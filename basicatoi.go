package piscine

func BasicAtoi(s string) int {
	var count, sum int
	for _, value := range s {
		if value >= '0' && value <= '9' {
			count = 0
			for i := '0'; 1 < value; i++ {
				count++
			}
			sum = sum*10 + count
		}
	}
	return sum
}

// Write a function that simulates the behaviour of the
// Atoi function in Go. Atoi transforms a number defined
// as a string in a number defined as an int.
// Atoi returns 0 if the string is not considered as a
// valid number. For this exercise only valid string will
//  be tested. They will only contain one or several digits as characters.
// For this exercise the handling of the signs + or -
// does not have to be taken into account.
// This function will only have to return the int.
// For this exercise the error return of Atoi is not required.
