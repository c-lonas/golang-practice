package main
import "fmt"


// Initialize Map Syntax:
// <map_nam> := map[<key_data_type>]<value_data_type>{key-value-pairs}






func main() {
	
	my_cool_map := map[string]int{"english": 85, "math": 80, "social_science": 90}

	english_grade := my_cool_map["english"]
	fmt.Println("My grade:", english_grade)


	// Add elements to the map simply by creating a new key and assigning a value
	my_cool_map["chemistry"] = 95
	fmt.Println(my_cool_map)


	// Delete elements from the map with the built-in delete method, supplying the name of the map and the key you want to delete
	delete(my_cool_map, "english")
	fmt.Println(my_cool_map)


	// Iterate over maps using 'range' again, in this case it returns the key and the value instead of the index and the element

	for key, value := range my_cool_map {
		fmt.Println(key, ":", value)
	}

	// You can also empty a map this way, by deleting each key/value pair from the map as you iterate through: `delete(my_cool_map, key)`

	// Alternatively, you can just re-initialize it as an empty map
	my_cool_map = make(map[string]int)
	fmt.Println(my_cool_map)

}