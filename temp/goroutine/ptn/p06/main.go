package main

import "log"

func main() {
	reqc := make(chan chan int)

	go func(cc chan chan int) {
		for {
			select {
			case resc := <-cc:
				log.Println("another goroutine receives a request")
				resc <- 42
			}
		}
	}(reqc)

	request := func() int {
		req := make(chan int)
		reqc <- req
		return <-req
	}

	log.Println("the main goroutine receives a response:", request())
}
