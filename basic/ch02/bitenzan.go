package main

import "fmt"

func main() {
	fmt.Println("5 / 2 = ", 5/2)
	fmt.Println("5 % 2 = ", 5%2)
	fmt.Println("3[00000011] & 6[00000110] = ", 3&6)
	fmt.Println("3[00000011] | 6[00000110] = ", 3|6)
	fmt.Println("3[00000011] &^ 6[00000110] = ", 3&^6)
	fmt.Println("3[00000011] << 1 = ", 3<<1)
	fmt.Println("3[00000011] >> 1 = ", 3>>1)
}
