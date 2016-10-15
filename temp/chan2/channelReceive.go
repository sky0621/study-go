package main

import "fmt"

type ChannelReceiver struct{}

type ChannelReceive interface {
	receiveChannel(v chan<- int)
}

// 渡されたチャネルから値を受信するメソッド
func (c *ChannelReceiver) receiveChannel(v <-chan int) {
	fmt.Println("call receiveChannel")
	for {
		val, ok := <-v
		if !ok {
			fmt.Println("ERROR!!!")
			break
		}
		fmt.Println(val)
	}
}
