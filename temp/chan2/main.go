package main

import "fmt"

func main() {
	fmt.Println("START")
	ch := make(chan int)

	sender := ChannelSender{max: 15}
	go sender.sendChannel(ch)

	// チャネル受信用の無限ループ
	for {
		// 受信用のチャネル ch から値を受け取る
		// ok は、チャネルがクローズされたかどうかの確認用
		value, ok := <-ch

		if !ok {
			break
		}

		fmt.Println(value)
	}

	fmt.Println("END")
}
