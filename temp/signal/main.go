package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer log.Println("Done !!!")

	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
	for _, trap := range trapSignals {
		log.Println(trap)
	}

	sigChan := make(chan os.Signal, 1) // OSからのシグナル受信用チャネル

	signal.Notify(sigChan, trapSignals...) // OSからシグナルがあった際にチャネルに通知する設定

	go doSomething()

	sig := <-sigChan // シグナル受信までブロック
	log.Println("Got Signal !!! ", sig)
}

func doSomething() {
	defer log.Println("done infinite loop.")
	for {
		time.Sleep(1 * time.Second)
	}
}
