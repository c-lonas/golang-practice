// Every Go program starts with a package declaration. We use main for executable progams.
package main

// Note that the package name in the import statement is surrounded with double quotes " "
import "fmt"


/*
You can write multi line comments between these delimiters
*/


// main function is a special function, and the entrypoint of executable programs. Every executable program must use this.
func main() {
	// The `dot` operator after fmt is the member access operator, used to access members of a package
	// In this case, we are accessing the Println function, a member of the fmt package
    fmt.Println("hello world")
}