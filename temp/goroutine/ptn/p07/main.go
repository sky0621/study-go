package main

import (
	"log"
	"math/rand"
	"time"
)

const goroutines = 3

func main() {
	c := make(chan int)

	for i := 0; i < goroutines; i++ {
		// 送信専用ゴルーチン
		go func(s chan<- int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			log.Println("処理完了")

			s <- 0 // 送信完了を伝えるだけなので値は何でもよい
		}(c)
	}

	for i := 0; i < goroutines; i++ {
		<-c // 送信したゴルーチンの分だけ、受信をウォッチ（＝処理完了を待機）
	}

	log.Println("すべて完了")
}
