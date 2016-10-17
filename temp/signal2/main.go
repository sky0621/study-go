package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MySignal struct {
	msg string
}

func (m MySignal) String() string {
	return m.msg
}

func (m MySignal) Signal() {}

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

	time.AfterFunc(10*time.Second,
		func() {
			sigChan <- MySignal{msg: "timed out !!"}
		})

	signal.Notify(sigChan, trapSignals...) // OSからシグナルがあった際にチャネルに通知する設定

	go doSomething()

	sig := <-sigChan // シグナル受信までブロック
	switch sig.(type) {
	case syscall.Signal:
		log.Println("Got Signal !!! ", sig)
	case MySignal:
		log.Println("App Original Signal !!!! ", sig)
	}
}

func doSomething() {
	defer log.Println("done infinite loop.")
	for {
		time.Sleep(1 * time.Second)
	}
}
