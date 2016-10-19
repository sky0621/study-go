package main

import "fmt"

func main() {
	m := "msg"
	foo(m)
}

func foo(m string) {
	s := "test"

	f := func() {
		fmt.Println(s + m)
	}
	f()
}
