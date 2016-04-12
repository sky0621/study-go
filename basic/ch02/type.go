package main

import "fmt"

func main(){
	type myInteger int
	var i myInteger = 123
	fmt.Println(i)
	i += 5
	fmt.Println(i)

	type myStruct struct {
		x int
		y int
		s string
	}
	var mys myStruct = new myStruct(3, 4, "test")
	fmt.Println(mys)
}
