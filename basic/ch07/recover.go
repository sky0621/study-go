package main

import "fmt"

func main(){
	f1(false)
	f1(true)
}

func f1(p bool){
	defer func(){
		fmt.Println("Defer func START")

		if err := recover(); err != nil {
			fmt.Println("Panic STOP", err)
		}

		fmt.Println("Defer func END")
	}()

	if p {
		fmt.Println("Panic occured")
	}
}
