package printer

import (
    "fmt"
    "os"
    "github.com/julianysavazzi/go-api/variables"
    "github.com/julianysavazzi/go-api/math"
)

const MAX int = 1000

func show() {
    fmt.Println("Hello World")
	fmt.Printf("My Name is %s\n", "Juliany")
	fmt.Printf("I am %d years old\n", 25)
	fmt.Printf("I am %f meters tall\n", 1.56)
	PrintMessage("Hello from printer package")
	variables.PrintVariables()

	n := 10

	// Flow control - choice what code lines to run
	if n < MAX {
		fmt.Printf("n %d is less than MAX %d \n", n, MAX)
	} else {
		fmt.Printf("n %d is greater than or equal to MAX %d \n", n, MAX)
	}

	expression := n < MAX
	switch expression {
	case true:
		fmt.Printf("n %d is less than MAX %d \n", n, MAX)
	case false:
		fmt.Printf("n %d is greater than or equal to MAX %d \n", n, MAX)
	default:
		fmt.Println("n is equal to MAX")
	}

	for n < (MAX - 950) {
		fmt.Printf("n %d is less than %d\n", n, (MAX - 950))

		if n % 2 == 0 {
			fmt.Println("n is even number\n")
		} else {
			fmt.Println("n is odd number\n")
		}

		n += 1
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i %d is less than 10\n", i)
		evenNumber := i % 2 == 0

		if evenNumber && i > 2 {
			continue
		}

		fmt.Println("here")

		if i == 5 {
			break
		}
	}

	// can be infinite foor, does not have while loop in GO
	for {
		n = n+1;
		fmt.Println(n)
		if n >= 10 {
			break
		}
	}

	sum := math.Add(1, 2)
	sub := math.Subtract(10, 6)
	mult := math.Multiply(10, 2)
	div, e := math.Divide(10, 2)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Subtract: %d\n", sub)
	fmt.Printf("Multiply: %d\n", mult)
	fmt.Printf("Divide: %d | error %s\n", div, e)
	fmt.Printf("Mod: %d\n", math.Mod(10, 2))

	result := math.Add(10, 100) // creation and attribution
	fmt.Printf("Result: %d\n", result)

	result = math.Subtract(result, 10) // reattribution
	fmt.Printf("Result: %d\n", result)

	// fmt.Println(math.DivideExample(10, 0))
	var err error
	result, err = math.Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Result: %d\n", result)

    args := os.Args // slice - list
    op := args[0]

    fmt.Println(args, op)
}