package main

import "fmt"

// int型をベースにmyType型を宣言
type myType int

func main() {
	var z myType = 12345
	z.show()
}

// myType型をレシーバにする関数を定義＝myTypeにメソッドを持たせた
func (v myType) show() {
	fmt.Println(v)
}
