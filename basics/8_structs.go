package main
import "fmt"

// Struct declaration syntax:
// type <struct_name> struct {
//		// list of fields
// }


type Student struct {
	name string
	id int
	gpa float64
}


// A method has a receiver, which takes a pointer to the instance of the associated struct
func (s *Student) bumpGrade() {
	s.gpa += 0.1
}


func main() {

	student_1 := Student {
		name: "josh",
		id: 0,
		gpa: 3.5,
	}

	// access fields of a struct using dot notation
	fmt.Println(student_1.name)
	fmt.Println(student_1.gpa)
	fmt.Println("Bumping gpa")
	student_1.bumpGrade()
	fmt.Println(student_1.gpa)



}