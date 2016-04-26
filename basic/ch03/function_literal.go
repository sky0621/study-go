package main

import "fmt"

func main() {
	func(i int) {
		fmt.Println(i * 2)
	}(10)

	f := func(s string) (i int) {
		if len(s) > 5 {
			i = 1
		} else {
			i = 0
		}
		return
	}
	i1 := f("abcde")
	fmt.Println(i1)
	i2 := f("abcdef")
	fmt.Println(i2)
}
