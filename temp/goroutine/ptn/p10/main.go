package main

import (
	"log"
	"time"
)

func main() {
	// 入室と退室を待ち受けるチャネルを持つルームを生成
	r := newRoom()
	go r.run()

	c1 := &client{name: "Taro", msg: "Hello!"}
	r.join <- c1
	r.left <- c1
	c2 := &client{name: "Hanako", msg: "Hi"}
	r.join <- c2
	r.left <- c2

	time.Sleep(5000)
	log.Println("main() fin.")
}

func newRoom() *room {
	return &room{
		join: make(chan *client),
		left: make(chan *client),
	}
}

type client struct {
	name string
	msg  string
}

type room struct {
	join chan *client
	left chan *client
}

func (r *room) run() {
	for {
		select {
		case c := <-r.join:
			log.Println("Join!!!")
			log.Println(c.name, ":", c.msg)
		case c := <-r.left:
			log.Println("Left...")
			log.Println(c.name, ":", c.msg)
		}
	}
}
