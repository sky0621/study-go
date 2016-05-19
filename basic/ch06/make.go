package main

import "fmt"

func main(){

	s1 := make([]int, 10, 20)	// 長さ10, キャパシティ20のスライス作成

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	fmt.Println()

	s2 := make([]float32, 5)

	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// スライスリテラル
	s := []int{1, 2, 3, 4}

	fmt.Println(s)

}
