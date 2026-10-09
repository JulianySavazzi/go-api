package math

import (
	"errors"
)

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func DivideExample(a int, b int) int {
	if b == 0 {
		panic("division by zero") // stop the application but is not captured
	}

	return a / b
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("division by zero") // this is the way GO capture the error
		return 0, err
	}

	return a / b, nil
}

func Mod(a int, b int) int {
	return a % b
}