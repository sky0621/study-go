package main

import "fmt"

// 計算インタフェース
type Calc interface {

	Calc(a int, b int) int

}

// 足し算型（あとで計算インタフェースを実装する）
type Add struct {

	// フィールドなし

}
func (x Add) Calc(a int, b int) int {

	return a + b

}

// 引き算型（あとで計算インタフェースを実装する）
type Sub struct {

	// フィールドなし

}
func (x Sub) Calc(a int, b int) int {

	return a - b

}

func main() {

	var ad Add
	var sb Sub

	fmt.Println(ad.Calc(3, 4))
	fmt.Println(sb.Calc(4, 3))

	var cal Calc
	cal = ad
	fmt.Println(cal.Calc(5, 3))
	cal = sb
	fmt.Println(cal.Calc(5, 3))

}

