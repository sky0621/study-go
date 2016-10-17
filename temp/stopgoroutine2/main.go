package main

import (
	"context"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Go 1.7 から標準パッケージ化されたコンテキストを使用
	ctx, cancel := context.WithCancel(context.Background())

	queue := make(chan string)

	for i := 0; i < 2; i++ { // ゴルーチンワーカーは２つ
		wg.Add(1)
		go fetchURL(ctx, queue)
	}

	queue <- "http://example.com/"
	queue <- "http://example.net/"
	queue <- "http://example.net/foo"
	queue <- "http://example.net/bar"

	cancel()  // コンテキスト（ctx）を終了させる
	wg.Wait() // 全ゴルーチンワーカーの終了を待機
}

func fetchURL(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			log.Println("worker exit")
			wg.Done()
			return
		case url := <-queue:
			log.Println("fetching: ", url)
			time.Sleep(3 * time.Second)
			log.Println("fetch fin:                        ", url)
		}
	}
}
