package main

import "fmt"

type Walker interface {
	Walk(s string) string
}

type Runner interface {
	Run(s string) string
}

type Exercise interface {
	Walker
	Runner
}

type Human struct {
	// Non
}
func (h Human) Walk(s string) string {
	return s + " walk!"
}
func (h Human) Run(s string) string {
	return s + " run!"
}

func main() {

	var h Human

	fmt.Println(h.Walk("Taro"))
	fmt.Println(h.Run("Taro"))

	var e Exercise
	e = h

	fmt.Println(e.Walk("Jiro"))
	fmt.Println(e.Run("Jiro"))

}

