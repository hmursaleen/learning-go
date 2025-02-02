package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//'errors' package is used to handle errors in Go. It provides a New function that you can use to create an error by passing in a string.

/*
In Go, a function whose name starts with a capital letter can be called by a
function not in the same package. This is known in Go as an exported name.
and a function whose name starts with a lowercase letter can't be called by
a function not in the same package.
*/
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name") //errors.New constructs a basic error value with the given error message.
	}

	// If a name was received, return a value that embeds the name
	//message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name) //
	/*In Go, the := operator is a shortcut for declaring and initializing a
	variable in one line (Go uses the value on the right to determine the variable's type).
	Taking the long way, you might have written this as:

	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	return message, nil //nil is used to indicate that the function executed successfully with no error.
	//Any Go function can return multiple values. For more, see Effective Go
}

// Hellos returns a map that associates each of the named people
// with a greeting message.

/*
In the last changes you'll make to your module's code, you'll add support for getting greetings for multiple people in one request. In other words, you'll handle a multiple-value input, then pair values in that input with a multiple-value output. To do this, you'll need to pass a set of names to a function that can return a greeting for each of them.

Note: This topic is part of a multi-part tutorial that begins with Create a Go module.
But there's a hitch. Changing the Hello function's parameter from a single name to a set of names would change the function's signature. If you had already published the example.com/greetings module and users had already written code calling Hello, that change would break their programs.

In this situation, a better choice is to write a new function with a different name. The new function will take multiple parameters. That preserves the old function for backward compatibility.
*/
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	//In Go, you initialize a map with the following syntax: make(map[key-type]value-type)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		//In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	//randomFormat starts with a lowercase letter, making it accessible only to code in its own package
	// A slice of message formats.
	formats := []string{
		//When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed.
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
