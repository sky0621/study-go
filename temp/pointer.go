package main

import "fmt"

func main() {
	var a *int
	var b = 2
	var c = 3

	fmt.Println("[実値]")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("############################")
	fmt.Println("[アドレス]")
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)
	fmt.Println("############################")
	fmt.Println("[[ a <- c ]]")
	a = &c
	fmt.Println("[実値]")
	fmt.Println(*a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("[アドレス]")
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)

}
