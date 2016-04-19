package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	arr := [...]int{0, 1, 2}
	for k := range arr {
		fmt.Println(k)
	}

	for l := 0; l < 5; l++ {
		if l == 1 {
			continue
		}
		if l == 4 {
			break
		}
		fmt.Println(l)
	}
}
