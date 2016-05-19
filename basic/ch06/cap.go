package main

import "fmt"

func main(){

	values := [...]int{0, 1, 2, 3, 4}

	s1 := values[1:3]

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := s1[1:4]	// キャパシティーまでスライスをスライス

	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := s1[1:3:4]

	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
}
