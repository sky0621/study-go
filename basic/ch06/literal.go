package main

import "fmt"

func main(){

	ary1 := [5]float32{}
	ary2 := [6]int{1, 2, 3, 4}	// 残りは０で初期化
	ary3 := [...]string{"One", "Two", "Three"}

	fmt.Println(ary1)
	fmt.Println(ary2)
	fmt.Println(ary3)
}
