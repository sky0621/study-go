package main

import "fmt"

func main() {

	// int型ポインタ
	var ip *int
	fmt.Println(" ip : ", ip)
	fmt.Println("&ip : ", &ip)
	fmt.Println("  ")

	var i int = 345
	fmt.Println("  i : ", i)
	fmt.Println(" &i : ", &i)

	// iのアドレスをポインタへ
	fmt.Println("  ")
	fmt.Println("<<< ip = &i >>>")
	fmt.Println("  ")
	ip = &i

	fmt.Println("  i : ", i)
	fmt.Println(" &i : ", &i)
	fmt.Println(" ip : ", ip)
	fmt.Println("&ip : ", &ip)
	fmt.Println("*ip : ", *ip)
	fmt.Println("  ")
	fmt.Println("  ")

	var x int = 6
	var y int = 7
	fmt.Println("<< before >>")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	double(x, &y)
	fmt.Println("<< after  >>")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

}

// yはポインタで受け取る
func double(x int, y *int) {
	x = x * 2
	*y = *y * 2
}
