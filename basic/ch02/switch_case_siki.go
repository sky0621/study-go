package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		switch true {
			case i % 2 == 0:
				fmt.Println("偶数", i)
			case i % 2 != 0:
				fmt.Println("奇数", i)
			default:
				fmt.Println(i)
		}
	}
}
