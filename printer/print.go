package printer // name of package must match the folder name

import "fmt"

func Printing() {
    printNumber(10)
    show()
}

// The public function must start with a capital letter
func PrintMessage(message string) {
	fmt.Println(message)
}

// The private function must start with a lowercase letter
func printNumber(number int) {
	fmt.Printf("Number: %d\n", number)
}