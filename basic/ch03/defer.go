package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer delay()
	fmt.Println("End")
}

func delay() {
	fmt.Println("delay")
}
