package main
import "fmt"

// In Go, an interface species a method set, and allows you to introduce modularity 
// interface is like a blueprint for a method set, by providing the function signature for each method. A little like traits in rust.
//
// Interface syntax: 
// type <interface_name> interface {
//	   // Method signature
// }


type pokemon interface {
	printType() string
	printLevel() string
}

type Bulbasaur struct {
	name string
	level int
	poketype string
}


func (b *Bulbasaur) printType() {
	fmt.Println(b.poketype)
}

func (b *Bulbasaur) printLevel() {
	fmt.Println(b.level)
}


func main() {

	bulby := Bulbasaur{
		name: "bulby",
		level: 5,
		poketype: "grass",
	}

	bulby.printLevel()
	bulby.printType()

}