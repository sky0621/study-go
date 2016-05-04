package main

import "fmt"

type myType int
func (v *myType) add(inc myType) myType {
	*v += inc
	return *v
}

func main() {

	var i myType
	fmt.Println(i.add(3))
	mv := i.add
	fmt.Println(mv(3))

}
