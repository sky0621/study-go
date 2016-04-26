package main

import "fmt"

func main() {

	// f1()実行結果をそのままf2()に渡せる！
	f2(f1())
}

func f1() (int, string, float32) {
	return 9, "Hi!", 3.14
}

// interface型はなんでも受け取れる型
func f2(i int, s string, c interface{}) {
	fmt.Println(i, s, c)
}
