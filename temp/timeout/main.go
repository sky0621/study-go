package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var sleep int
	if len(os.Args) != 2 || os.Args[1] == "" {
		sleep = 2
	} else {
		sleep, _ = strconv.Atoi(os.Args[1])
	}
	timer := time.NewTimer(time.Duration(sleep) * time.Second)

	done := make(chan int)

	go func() {
		done <- doSomething()
	}()

	select {
	case <-timer.C:
		log.Println("timeout REACHED !!!")
	case r := <-done:
		log.Println("doSomething !!! ", r)
	}
}

func doSomething() int {
	log.Println("[START]")
	time.Sleep(3 * time.Second)
	log.Println("[END]")
	return 0
}
