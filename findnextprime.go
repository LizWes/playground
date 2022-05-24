package piscine

func FindNextPrime(index int) int {
	if index < 0 {
		return -1
	}
	if index <= 1 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}

// Write a function that returns the first prime number
// that is equal or superior to the int passed as parameter.
// The function must be optimized in order to avoid time-outs with the tester.
// (We consider that only positive numbers can be prime numbers)
