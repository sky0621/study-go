package main

import "log"

func main() {
	resc := make(chan int, 8)

	go func(resc chan int) {
		for i := 0; i < 42; i++ {
			log.Println("another goroutine is sending:", i)
			resc <- i
		}
		close(resc)
	}(resc)

	for i := range resc {
		log.Println("the main goroutine receives:", i)
	}
}
