package main
import "fmt"


func main() {

	// Three different ways to initialize Arrays: "Regular", shorthand, ellipses
	// Index into them like normal
	var toppings [2]string = [2]string{"Cheese", "Tomatoes"}
	fmt.Println(toppings)

	nums := [3]int{23, 46, 85}
	fmt.Println(nums[0])

	countries := [...]string{"France", "Japan", "Belgium"}

	// Iterate through an array
	for i := 0 ; i < len(countries) ; i ++ {
		fmt.Println(countries[i])
	}
	

	// Can also iterate through array using `range` keyword, which is like enumerate in python

	pokemon := [...]string{"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon", "Charizard", "Squirtle", "Wartortle", "Blastoise"}

	for index, element := range pokemon {
		fmt.Println(element, "is pokemon #", index)
	}



	// Multidimensional arrays ("Array of arrays")
	
	// 3 inner arrays, each of which hold 2 elements
	arr := [3][2]int{ {2, 5}, {8, 11}, {5, 100} }

	// Access the first array with the first index, and then access the second element of that array with the next index
	fmt.Println(arr[0][1])
}