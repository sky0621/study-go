package main

import "fmt"
import "unicode/utf8"

func main(){
	// 6byte
	var en string = "golang"
	// utf-8のためjaは1文字3byte
	var ja string = "Go言語"
	fmt.Println(en, " len:", len(en))
	fmt.Println(ja, " len:", len(ja))

	fmt.Println(ja, " 文字列長：", utf8.RuneCountInString(ja))
}
