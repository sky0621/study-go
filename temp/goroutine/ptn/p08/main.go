package main

import (
	"log"
	"os"
)

const goroutines = 10

func main() {
	counter := make(chan int, 1) // １個のみに限定したカウントアップ値保持用チャネル

	for i := 0; i < goroutines; i++ {
		go func(cntr chan int) {
			value := <-cntr
			value++
			log.Println("counter: ", value)

			if value == goroutines {
				os.Exit(0)
			}

			cntr <- value
		}(counter)
	}

	// カウンターに初期値を送信
	counter <- 0

	for {
		// 無限ループ
	}
}
