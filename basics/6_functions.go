package main
import "fmt"

// Note: a higher order function is a function that takes another function as an argument, or returns an argument



func main() {

	fmt.Println(adder(1, 2))
	fmt.Println(sumAndDiff(5, 4))
	fmt.Println(sumAndDiffNamed(3, 7))
	fmt.Println(variadicSum(4, 7, 22, 5))

	n := 5
	fmt.Println("Factorial of", n, ":", factorial(n))

	// Anonymous (Lambda) functions:
	x := func(a int, b int ) int {
		return a * b
	}

	// Here we stored a function in variable `x` without ever naming the function itself

	// Now we can use the function we stored in x	
	fmt.Println(x(2, 6))


	// Alternatively, you can just create an anonymous function and use it directly, and store the output, not the function itself, in a variable
	y := func(a int, b int) int {
		return a * b
	}(20, 5)
	fmt.Println(y)

}



// Basic function syntax
func adder(a int, b int) int {
	return a + b
}


// Return multiple values
func sumAndDiff(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}


// Can also do named return values- note that if we initialize the return values in this way,
// we can assign the values with only a '=' instead of ':='.
// Also, you no longer need to specify the variable names in the return statement.
func sumAndDiffNamed(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}


// Variadic functions:
// Functions that can accept varying numbers of arguments of the same type
// To declare a variadic function, the type of the final parameter is preceded by `...`
func variadicSum(numbers ...int) (sum int) {
	sum = 0
	for _, value := range numbers {
		sum += value
	}
	return 
}

// Recursive functions:
// Functions that directly or indirectly calls itself
// Recursive function will keep calling itself until it reaches a base case

func factorial(num int) int {

	// First, define the base case
	if num == 1 {
		return 1
	}

	// Then, define the 'recurse' case
	return num * factorial(num - 1)
}




