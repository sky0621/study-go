package main

import "fmt"

func main() {
	ans := isPlus(3, 2)
	fmt.Println(ans)
	ans = isPlus(2, 3)
	fmt.Println(ans)

	fmt.Println("  ")

	x, y := 5, 6
	fmt.Println(x, y)
	add, sub, mul, div := calc(x, y)
	fmt.Println(add, sub, mul, div)
}

func isPlus(a int, b int) string {
	if a - b > 0 {
		return "Plus!"
	} else {
		return "Minus!"
	}
}

func calc(x int, y int) (int, int, int, float32) {
	return x + y, x - y, x * y, float32(x) / float32(y)
}
