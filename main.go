/**
* Package main provides entrypoint for the application.
* Package (folders) -> contains many files;
* Files (in folder) -> contains many functions;
* Function (in file) -> belongs to a package;
* For run the application, uses go run filename.go;
* Go is a strongly typed programming language;
* 
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Printf("My Name is %s\n", "Juliany")
	fmt.Printf("I am %d years old\n", 25)
	fmt.Printf("I am %f meters tall\n", 1.56)
}