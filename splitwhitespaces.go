package piscine

func SplitWhiteSpaces(s string) []string {
	TextToString := ""
	newTextSlice := []string{} //the new text slice
	for index, value := range s {
		if (index == LenghtOf(s)-1) && (string(value) != " ") && (string(value) != "\t") && (string(value) != "\n") {
			TextToString += string(value)
			newTextSlice = append(newTextSlice, TextToString)
		} else if string(value) != " " && string(value) != "\t" && string(value) != "\n" { //if
			TextToString += string(value)
		} else {
			if len(TextToString) >= 1 {
				newTextSlice = append(newTextSlice, TextToString)
			}
			TextToString = ""
		}
	}
	return newTextSlice
}

func LenghtOf(d string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}

// Write a function that separates the words of
//a string and puts them in a string slice.
//The separators are spaces, tabs and newlines.
