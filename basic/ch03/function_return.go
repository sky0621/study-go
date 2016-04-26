package main

import "fmt"

func main() {
	ret := add(3, 4)
	fmt.Println(ret)
}

func add(x int, y int) (xpy int) {
	xpy = x + y
	return
}
