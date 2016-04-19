package main

import "fmt"

func main() {
	const G float32 = 9.80665
	const G2 = G * 2
	fmt.Println(G2)

	const (
		ZERO = iota
		ONE
		TWO
		THREE
	)
	fmt.Println(ZERO, ONE, TWO, THREE)

	const (
		bit0 = 1 << iota
		bit1
		bit2
	)
	fmt.Println(bit0, bit1, bit2)
}
