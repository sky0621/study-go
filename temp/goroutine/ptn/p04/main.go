package main

import (
	"log"
	"time"
)

func main() {
	log.Println("========== buffer[0], sleep[0] ==========")
	resc := make(chan int)
	go sendInt(resc)
	show(resc, 0)

	log.Println("")
	log.Println("========== buffer[5], sleep[0] ==========")
	resc2 := make(chan int, 5)
	go sendInt(resc2)
	show(resc2, 0)

	log.Println("")
	log.Println("========== buffer[0], sleep[3] ==========")
	resc3 := make(chan int)
	go sendInt(resc3)
	show(resc3, 3000)

	log.Println("")
	log.Println("========== buffer[5], sleep[3] ==========")
	resc4 := make(chan int, 5)
	go sendInt(resc4)
	show(resc4, 3000)

	log.Println("")
	log.Println("========== buffer[10], sleep[0] ==========")
	resc5 := make(chan int, 10)
	go sendInt(resc5)
	show(resc5, 0)

	log.Println("")
	log.Println("========== buffer[10], sleep[3] ==========")
	resc6 := make(chan int, 10)
	go sendInt(resc6)
	show(resc6, 3000)
}

func sendInt(c chan int) {
	for i := 0; i < 11; i++ {
		log.Println("another goroutine is sending:", i)
		c <- i
	}
	close(c)
}

func show(c chan int, sec time.Duration) {
	time.Sleep(sec)
	for i := range c {
		log.Println("the main goroutine receives:", i)
	}
}
