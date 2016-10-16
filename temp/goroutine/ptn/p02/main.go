package main

import "github.com/sky0621/study-go/temp/fslog"

var lg = fslog.FsLogger{FilePathStr: "./fslog.log"}

func main() {
	resc := make(chan int)

	go func(c chan int) {
		lg.Println("another goroutine")
		c <- 42
	}(resc)

	lg.Println("the main goroutine, received value:", <-resc)
}
