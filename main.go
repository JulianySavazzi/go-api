/**
* Package main provides entrypoint for the application.
* Package (folders) -> contains many files;
* Files (in folder) -> contains many functions;
* Function (in file) -> belongs to a package;
* Go is a strongly typed programming language;
* Go is a compiled programming language;
* code -> compiler -> machine code -> run file;
* For run the application, uses go run filename.go;
* For build the application, uses go build filename.go -> generate an executable file;
* GOOS -> Windows, Linux, MacOS;
* GOARCH -> 64 bits, 32 bits, arm;
* go version -> show go version and architecture of system where the command is executed;
* Example for build the application for Windows: GOOS=windows GOARCH=amd64 go build filename.go;
* CI: compilation;
* CD: run the binary file on production machine;
* GO binary file: a GO binary is self-contained, all the system libraries required to run the app are included within it;
* go mod init github.com/username/repo-name -> create a module (a collection of packages - required dependencies);
*/
package main

import (
	"fmt"
	"github.com/julianysavazzi/go-api/printer"
	"github.com/julianysavazzi/go-api/variables"
)

const MAX int = 1000

func main() {
	fmt.Println("Hello World")
	fmt.Printf("My Name is %s\n", "Juliany")
	fmt.Printf("I am %d years old\n", 25)
	fmt.Printf("I am %f meters tall\n", 1.56)
	printer.PrintMessage("Hello from printer package")
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
}