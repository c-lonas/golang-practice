package main
import "fmt"

// A pointer is a variable that points to the memory address of another variable


// Address operator:       &
// Dereference operator:   *




func main() {


	///// Part I

	// Initialize a variable
	x := 42
	fmt.Println(x)

	// Print hexidecimal memory address of that variable
	fmt.Println(&x)

	// The hexidecimal memory address will change with each execution, so for this illustration we'll store the address in `y`
	y := &x
	fmt.Println(y)

	// This is the equivalent of writing `fmt.Println(*0xc00001c0a8)`, but again, the address changes every time, hence y
	fmt.Println(*y)


	// You can also use the dereference operator to assign a new value to that memory location (printing `x` here will return 12)
	*y = 12
	fmt.Println(x)

}