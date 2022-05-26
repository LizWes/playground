# playground
Things I did a magical month in Åland

There is a comment in every function and program indicating the task given.

Here is the list of functions and programs I did and short explinations:

**printalphabet/main.go** A program that prints the latin alphabet as a string
 
**printreversealphabet/main.go** A program that prints the latin alphabet as a string in reversorder

**printdigits/main.go** A program that prints decimal digits from 0 to 9 as runes

**isnegative.go** func IsNegative(number int)
// A function that prints T if its negative or F if it is not

**printcomb.go** func PrintComb()
// A function that returns thre number combinations in ascending order where firstDigit < secondDigit < thirdDigit

**pointone.go** func PointOne(n *int) 
// A function that changes the value of an int trough a pointer to 1

**ultimatepointone.go** func UltimatePointOne(n ***int) 
//A function that changes the value of an int trough three pointers to 1

**divmod.go** DivMod(a int, b int, div *int, mod *int) 
// A function that takes two ints and returns the division result of them to one pre determinde variable and the remainder of the divison to another pre determind variable

**ultimatedivmod.go** UltimateDivMod(a *int, b *int) 
//A function that takes two ints and returns the division result of them to the first variable and the remainder of the divison to second variable

**printstr.go** func PrintStr(s string) 
// A function that prints a string one character at a time as a rune 

**strlen.go** func StrLen(s string) int 
// A function that counts the lenght of a string and returns it as a int

**strrev.go** func StrRev(s string) string
// A function that reverses a string

**iterativefactorial.go** func IterativeFactorial(nb int) int 
// A function that returns the factorial of the number, overflows or 0 value will return 0

**recursivefactorial.go** func RecursiveFactorial(nb int) int 
// A function that returns the factorial of the number, overflows or 0 value will return 0

**iterativepower.go** func IterativePower(nb int, power int) int 
// A function that returns the number times the factorial of power

**recursivepower.go** func RecursivePower(nb int, power int) int 
// A function that returns the number times the factorial of power

**fibonacci.go** func Fibonacci(index int) int
// A function that returns the number of the at set index from Fibonacci sequence

**sqrt.go** func Sqrt(nb int) int
// A function that returns the square root of a number, if it's a whole number

**isprime.go** func IsPrime(nb int) bool
// A function that checks if the number is a prime number

**findnextprime.go** func FindNextPrime(index int) int 
// A function that returns the next prime number or the number passed as index if it's a prime number

**firstrune.go** func FirstRune(s string) rune 
// A function that converts a string to a slice that returns the first caracter and returns it as a rune

**lastrune.go** func LastRune(s string) rune
// A function that converts a string to a slice that returns the last caracter and returns it as a rune

**nrune.go**  func NRune(str string, num int) rune
// A function that converts a string to a slice and returns a character using the num as index for the slice

**compare.go** func Compare(a, b string) int 
// A function that compares two strings and returns 0 it they are equeal to each other, -1 if a is smaler than b and 1 if a is bigger than b

**alphacount.go** func AlphaCount(s string) int 
// counts the amount of letters in a string and returns the count

**index.go** func Index(s string, toFind string) int
// A function that looks if the string toFind is found in string s and returns -1 if its not found and 1 if its found

**concat.go** (str1 string, str2 string) string 
// A function that adds to strings togheter and returns the result


