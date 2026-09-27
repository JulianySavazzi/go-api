package printer // name of package must match the folder name

import "fmt"

// The public function must start with a capital letter
func PrintMessage(message string) {
	fmt.Println(message)
	printNumber(10)
}

// The private function must start with a lowercase letter
func printNumber(number int) {
	fmt.Printf("Number: %d\n", number)
}