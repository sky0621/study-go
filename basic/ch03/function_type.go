package main

import "fmt"

func main(){
	var f func(int, int) int

	f = func(a int, b int) int {
		return a + b
	}

	fmt.Println(f(9, 12))

	f = multi
	fmt.Println(f(8, 7))
}

func multi(x int, y int) int {
	return x * y
}

func init() {
	fmt.Println("Initial!")
}
