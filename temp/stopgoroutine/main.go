package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ { // ゴルーチンワーカーは２つ
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "http://example.com/"
	queue <- "http://example.net/"
	queue <- "http://example.net/foo"
	queue <- "http://example.net/bar"

	close(queue) // チャネルは一度しかクローズできない。クローズ済みチャネルに送信するとパニック！
	wg.Wait()    // 全ゴルーチンワーカーの終了を待機
}

func fetchURL(queue chan string) {
	for {
		url, more := <-queue
		if more {
			log.Println("fetching: ", url)
			time.Sleep(3 * time.Second)
			log.Println("fetch fin:                        ", url)
		} else {
			log.Println("worker exit")
			wg.Done()
			return
		}
	}
}
