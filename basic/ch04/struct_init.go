package main

import "fmt"

type Person struct {
	name string
	age int
}

func main(){
	var p1 Person
	p1.name = "Taro"
	p1.age = 37
	fmt.Println(p1)

	p2 := Person{age: 25, name: "Jiro"}
	fmt.Println(p2)
	
	p3 := Person{"Saburo", 19}
	fmt.Println(p3)

	p4 := &Person{"Siro", 12}
	fmt.Println(p4)

}
