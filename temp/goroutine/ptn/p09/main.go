package main

import (
	"log"
	"os"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	// 各チャネルへの送信ゴルーチン
	go func() {
		for i := 0; i < 10; i++ {
			// 送信時もselect使用
			select {
			case ch1 <- i:
				log.Println("ch1への送信完了！")
			case ch2 <- "test":
				log.Println("ch2への送信完了！")
			}
		}
		os.Exit(0)
	}()

	for {
		select {
		case val := <-ch1:
			log.Println("ch1からの受信完了：", val)
		case text := <-ch2:
			log.Println("ch2からの受信完了：", text)
		}
	}
}
