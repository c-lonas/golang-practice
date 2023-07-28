package main

import "fmt"


func main() {

	// First method for declaring variables
	var my_string string = "Hello World"
	var my_int int = 100
	var my_float float64 = 8125.456782
	var my_bool bool = false


	fmt.Println(my_string, "my number is", my_int)
	fmt.Printf("Here are my other variables: %v , %v \n", my_bool, my_float )


	// Second method is to initialize the variable first, and assign it a value later

	var pokemon string
	pokemon = "bulbasaur"
	fmt.Println("My pokemon is", pokemon)



	// A third method is to declare multiple variables of the same type on a single line

	var my_special_num, your_special_num int = 100, 500
	fmt.Println("Special numbers:", my_special_num, your_special_num)


	// To declare multiple variables together if they are of different types, you can ise this syntax:

	var (
		lightswitch_on bool = true
		lightswitch_color string = "white"
	)
	fmt.Println("The", lightswitch_color, "lightswitch is on:", lightswitch_on )



	// There is another shorthand variable declaration method called 'short variable declaration'
	// With this method, you don't need to use the var keyword or explicitly assign a type
	s:= "Hello, World!"
	fmt.Println(s)



	// you can declare a constant with the `const` keyword - these are immutable
	// Constants have to be declared and assigned at the same time, you can't initialize it and assign later like you can with a regular variable

	// untyped constant
	const age = 12

	//
	const guitar_type string = "Stratocaster"
}