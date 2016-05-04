package main

import "fmt"

func main() {
	var Book struct {
		id int
		name string
	}

	Book.id = 345
	Book.name = "Sample Book"

	fmt.Println(Book)
}
