package main

import (
	"log"
	"net/http"
)

var urls = []string{
	"http://example.com/",
	"http://example.net/",
	"http://example.org/",
}

func main() {
	log.Println("[main() start]")
	// HTTPステータスをゴルーチン間でやりとりするためのチャネル
	statusChannel := make(chan string)
	defer close(statusChannel)

	for _, url := range urls {
		go doGet(url, statusChannel)
	}

	for i := 0; i < len(urls); i++ {
		log.Println("[doGet END]  " + urls[i] + ": " + <-statusChannel) // チャネルからHTTPステータスを取得
	}
	log.Println("[main() end]")
}

// 送信専用のチャネルとして sc 定義
func doGet(url string, sc chan<- string) {
	log.Println("[doGet START]" + url)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	sc <- res.Status
}
