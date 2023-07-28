package main
import "fmt"

func main() {



	// Conditionals

	var s string = "pizza"
	// If-else block
	if s == "pizza" {
		fmt.Println(s)

	} else {
		fmt.Println("Still hungry")
	}

	// Switch Statement

	var pokemon string = "squirtle"
	switch pokemon {
		case "squirtle":
			fmt.Println("I choose you, Squirtle!")
		case "charmander":
			fmt.Println("I choose you, Charmander!")
		case "bulbasaur":
			fmt.Println("I choose you, Bulbasaur!")
		default:
			fmt.Println("The early bird gets the worm- or, in this case, the Pokemon")
	}

	// Fallthrough can be used to force the switch case to force execution flow to 'fall through' the successive case block

	var result string = "fail"
	switch result {
		case "fail":
			fmt.Println("You lost")
			fallthrough
		case "success":
			fmt.Println("Actually, you win!")
			fallthrough
		default:
			fmt.Println("Thanks for playing!")
	}



	// For Loop
	// initialize variable, condition, post
	for i := 0; i <= 5; i++ {
			fmt.Println(i * i)
			if i == 3 {
					break
			}
	}
}