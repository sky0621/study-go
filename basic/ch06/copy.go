package main

import "fmt"

func main(){

	dest := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	count := copy(dest[2:], src)
	fmt.Println(count)

	fmt.Println(dest)

}
