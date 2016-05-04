package main

import "fmt"

type base struct {
	i int
}
func (b base) show() {
	fmt.Println("show!")
}

type target struct {
	base
}

func main() {
	var x target
	x.show()
	fmt.Println(x.i)
}
