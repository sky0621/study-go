package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan bool)
	m := make(map[string]string)

	var mux sync.Mutex

	go func() {
		mux.Lock()
		m["1"] = "a"
		mux.Unlock()
		ch <- true
	}()

	mux.Lock()
	m["2"] = "b"
	mux.Unlock()

	<-ch // ch はバッファ無しで生成しているので go func() 内での ch への送信を待機する。
	// メインと別ゴルーチンの両方の処理が終わってから出力
	for k, v := range m {
		fmt.Println(k, v)
	}
}
