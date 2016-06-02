package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("main:START")

	fmt.Println("call test()")
	test()

	fmt.Println("call go test()")
	go test()

	time.Sleep(3 * time.Second)

	fmt.Println("main:END")
}

func test(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		time.Sleep(1 * time.Second)
	}
}
