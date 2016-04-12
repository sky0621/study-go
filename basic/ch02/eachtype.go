package main

import "fmt"

func main(){
	// bool
	var b bool
	fmt.Println(b)
	b = true
	fmt.Println(b)
	b = true || false
	fmt.Println(b)

	// number
	var i int
	fmt.Println(i)
	i = 12345
	fmt.Println(i)
	var i2 int64
	fmt.Println(i2)
	i2 = int64(i)
	fmt.Println(i2)

	// string
	var s string
	fmt.Println(s)
	s = "あ"
	fmt.Println(s)
	s = s + "い"
	fmt.Println(s)
	s += "う"
	fmt.Println(s)
}
