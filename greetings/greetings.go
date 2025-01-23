package greetings

import "fmt"

/*In Go, a function whose name starts with a capital letter can be called by a
function not in the same package. This is known in Go as an exported name.
*/
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	/*In Go, the := operator is a shortcut for declaring and initializing a
	variable in one line (Go uses the value on the right to determine the variable's type).
	Taking the long way, you might have written this as:

	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	return message
}
