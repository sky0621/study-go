package main

import "log"

func main() {

	// 初回実行後、しばらくはプロセスが生き残ってる？（"another goroutine"のログがしばらく出なくなる）
	go func() {
		log.Println("another goroutine")
	}()

	log.Println("the main goroutine")
}
