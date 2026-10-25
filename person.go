package main

// data structure
type Person struct {
    Name string
    Surname string
    Age int
    City string
}

// Is a Person method, must be call by Person instance
func (p Person) Greeting() string {
    return "Hello, my name is " + p.Name + "!"
}

// structs for reference value
func (p *Person) IncrementAge() {
    // it uses a *Person pointer, so it modifies the original Person instance instead of working with a copy
    p.Age += 1
}

// In Go, you do not need to explicitly define a constructor as you would in other object-oriented languages.
// Instead, you can create functions that return a new instance of a struct, serving as a form of constructor.
func NewPerson( name string, surname string, age int, city string) Person {
    return Person {
        Name: name,
        Surname: surname,
        Age: age,
        City: city,
    }
}