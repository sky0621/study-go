package main

import "fmt"

func main() {
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	var (
		i int = 456
		s1, s2 = "A", "B"
	)
	fmt.Println(i, s1, s2)

	t := 789
	fmt.Println(t)
}
