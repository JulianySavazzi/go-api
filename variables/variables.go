package variables

import "fmt"

// global scope in package

const Pi float64 = 3.14159 // declaration of a constant (can't be changed, may be declared and attributed)

const timeToSleepInSeconds = 500

var test string

// strongly and staticly typed programming language
func PrintVariables() {
	// function scope

	test = "teste"
	fmt.Println(test)

	x := "Hello World" // explicit declaration and attribution of value, static type of x is string
	fmt.Println(x) // references to x

	var y string // creation and declaration of a variable
	y = "Hello World" // attribution of value
	fmt.Println(y) // references to y

	var a, b, c int = 1, 2, 3 // declaration and atribution of values to multiple variables
	fmt.Println(a, b, c) // references to variables

	d, e := 5, 10.5 // short declaration and attribution of values to multiple variables
	fmt.Println(d, e) // references to variables

	var f, g int // declaration of multiple variables
	fmt.Println(f, g)

	g = 2
	fmt.Println(f) // default value for int variables is 0
	fmt.Println(g)

	var h bool
	h = true
	fmt.Println(h)

	var i rune
	i = 'a'
	fmt.Println(i)

	fmt.Println(Pi)

	fmt.Println(timeToSleepInSeconds)

	showString(test)
	showString("Oi")

}

func showString(text string) {
	fmt.Println(text)
}