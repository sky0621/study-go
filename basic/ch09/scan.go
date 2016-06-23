package main

import "fmt"

func main(){
	var fname, name string
	fmt.Println("Name?")
	fmt.Scanln(&fname, &name)
	fmt.Println("Age?")
	var age int
	fmt.Scanf("%d", &age)
	fmt.Printf("Name:%s %s Age:%d\n", fname, name, age)
}
