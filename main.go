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

func main() {
	fmt.Println("Hello World")
	fmt.Printf("My Name is %s\n", "Juliany")
	fmt.Printf("I am %d years old\n", 25)
	fmt.Printf("I am %f meters tall\n", 1.56)
	printer.PrintMessage("Hello from printer package")
	variables.PrintVariables()
}