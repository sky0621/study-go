package main

import "fmt"

func main() {
	fmt.Println("Hello")

	no := 3
	if no % 2 == 0 {
		fmt.Println("gusu")
	} else {
		fmt.Println("kisu")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
