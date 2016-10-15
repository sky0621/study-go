package main

import "fmt"

type ChannelSender struct {
	max int // 送信する最大値
}

type ChannelSend interface {
	sendChannel(v chan<- int)
}

// 渡されたチャネルに値を送るメソッド
func (c *ChannelSender) sendChannel(v chan<- int) {
	fmt.Println("call sendChannel(chan)")
	for i := 0; i < c.max; i++ {
		v <- i * i
	}
	close(v)
}
