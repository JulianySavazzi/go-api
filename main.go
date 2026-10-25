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
	"os"
	"github.com/julianysavazzi/go-api/printer"
	"github.com/julianysavazzi/go-api/math"
)

func main() {
    // this func must contains all essential basic functions for learning go
//     printer.Printing()

    args := os.Args // slice -> list of values

//     op := os.Args[1]
//     n1 := os.Args[2]
//     n2 := os.Args[3]

    printer.PrintMessage("Type an operation, and two numbers:")
    fmt.Printf("Arguments %d \n", args)
//     fmt.Printf("Executing a operation %s with values %d and %d\n", op, n1, n2)

    division, err := math.Divide(10, 2)
    fmt.Printf("Division %d, error %s\n", division, err)

// go mod tidy:
// for update go.mod and go.sum with dependencies (external libs), used in your code and
// to remove unused dependencies from your go.mod and go.sum

// go get:
// for make download by external specific dependencies and update go.mod and go.sum

// go mod init:
// for create go.mod for manage dependencies and modules

    // using struct
    var p Person
    p.Name = "Juliany"
    p.Surname = "Savazzi"
    p.Age = 26
    p.City = "Votuporanga"

    fmt.Printf("Person %+v\n", p)

    p2 := Person{ "Matheus", "Paula", 28, "Ilha Solteira"}

    fmt.Printf("Person %+v\n", p2)

    fmt.Println(p.Greeting())
    helloMath := p2.Greeting()
    fmt.Println(helloMath)

    p3 := NewPerson("Bruno", "Savazzi", 22, "Votuporanga")
    fmt.Println(p3)

    p3.IncrementAge()
    fmt.Println(p3.Age)
}