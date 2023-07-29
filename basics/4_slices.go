package main
import "fmt"


// Slices are continuous segments of an array
// Slices can have elements added or removed

func main() {

	arr := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	my_first_slice := arr[2:8]
	fmt.Println(my_first_slice)

	// Can also create slices from other slices
	my_second_slice := my_first_slice[1:4]
	fmt.Println(my_second_slice)


	// If you change an element in the array
	arr[2] = 0

	// The change will also be reflected in the slice
	fmt.Println("Changed 3rd element to 0")
	fmt.Println(arr)
	fmt.Println(my_first_slice)


	// You can append elements to the slice using the built-in 'append' method
	my_new_slice := append(my_first_slice, 1000)
	fmt.Println(my_new_slice)

}