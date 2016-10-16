package main

import "log"

func main() {
	resc := make(chan int)

	go func(c chan int) {
		for i := 0; i < 42; i++ {
			log.Println("another goroutine is sending:", i)
			c <- i
		}
		close(c)
	}(resc)

	for i := range resc {
		log.Println("the main goroutine receives:", i)
	}
}
