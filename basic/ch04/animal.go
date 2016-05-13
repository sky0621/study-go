package main

import "fmt"

type Animal struct {
	name string
	kind string
}

func main() {

	var cat Animal
	cat.name = "Mike"
	cat.kind = "cat"
	fmt.Println(cat)

	dog := Animal{"Pochi", "dog"}
	fmt.Println(dog)

}
