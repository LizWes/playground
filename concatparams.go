package piscine

func ConcatParams(args []string) string {
	var index int
	var newString string
	for i := range args {
		index = i
	}
	for i := 0; i <= index; i++ {
		if i != index {
			newString = newString + args[i] + "\n"
		} else {
			newString = newString + args[i]
		}
	}
	return newString
}

// Write a function that takes the arguments received
// in parameters and returns them as a string. The string
// is the result of all the arguments concatenated with a newline (\n) between.
