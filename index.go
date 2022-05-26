package piscine

func Lenght(str string) int {
	counter := 0
	for range str {
		counter++
	}
	return counter
}

func Index(s string, toFind string) int {
	if toFind == "" { //if there is no value ToFind
		return 0
	}
	converteds := []rune(s)
	convertedToFind := []rune(toFind)
	lenght := Lenght(toFind)
	ind := -1 // Default value to be changed if the ToFind string is found in the string
	counter := 0
	for i := 0; i < Lenght(s); i++ { // loop comparing every part of the string
		if converteds[i] == convertedToFind[counter] { // if the character matches the caracter of ToFind that the counter currently is on
			if counter == 0 { // If this is the first character ind gets the index value of the chatacter
				ind = i
			}
			counter++              // in all instances of that the characters match the counter will increase
			if counter == lenght { //if the the counter is equal to the lenght of the the string ToFind Index is returnd
				return ind
			}
		} else {
			if counter > 0 { // If the caracters are not equal, ind and counter is reset, if counter is greater than zero the for loop "goes back" one value
				i--
			}
			counter = 0
			ind = -1
		}
	}
	return ind
}

//Write a function that behaves like the Index function.
