package main

import (
	"time"
	"fmt"
	"os"
)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine")
		os.Exit(0)
	}()

	// 無限ループ
	for {}
}
