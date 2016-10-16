package main

import "log"

func main() {
	ch := make(chan int, 3)

	ch <- 0

	log.Println("cap: ", cap(ch))
	log.Println("len: ", len(ch))

	ch <- 1

	log.Println("cap: ", cap(ch))
	log.Println("len: ", len(ch))

	ch <- 2

	log.Println("cap: ", cap(ch))
	log.Println("len: ", len(ch))

	log.Println("---------- FULL ----------")

	ch <- 3

	log.Println("cap: ", cap(ch))
	log.Println("len: ", len(ch))

}
