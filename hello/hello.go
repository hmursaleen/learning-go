package main

/*Declare a main package (a package is a way to group functions,
and it's made up of all the files in the same directory).
The main package is special—it's the package that tells Go to create an executable.
*/

import (
	"fmt"
	"log" //log package is used to log messages in Go.

	"example.com/greetings"
)

// Import the fmt package, which is Go's standard input/output library. fmt is short for format.

/*
You can use the pkg.go.dev site to find published modules whose packages
have functions you can use in your own code. Packages are published in
modules -- like rsc.io/quote -- where others can use them.

When you import a module from online(pkg.go.dev) Go will add the quote module
as a requirement, as well as a go.sum file for use in authenticating the module.

and you install the requirements by running following command, like we use
pip install -r requirements.txt in python:

$ go mod tidy



In a module, you collect one or more related packages for a discrete and useful
set of functions. For example, you might create a module with packages that
have functions for doing financial analysis so that others writing financial
applications can use your work.

Go code is grouped into packages, and packages are grouped into modules.
Your module specifies dependencies needed to run your code, including the
Go version and the set of other modules it requires.





*/
//A main function executes by default when you run the main package.
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	// Change the Hello argument from Gladys’s name to an empty string, so you can try out your error-handling code.
	//message, err := greetings.Hello("Mursaleen") // Assign both of the Hello return values, including the error
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err) //log.Fatal function call will print the error and stop the program.
	}

	// If no error was returned, print the returned message
	// to the console.
	//message := greetings.Hello("Gladys")
	fmt.Println(messages)
}
