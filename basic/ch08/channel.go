package main

import "fmt"
import "time"

func main(){
	// チャネル生成
	c := make(chan int)

	// 送信専用チャネル
	go func(s chan<- int){
		for i := 0; i < 5; i++ {
			// チャネル s へ0～4を順番に送信
			s <- i
			time.Sleep(1 * time.Second)
		}
		// チャネルのクローズ
		close(s)
	}(c)

	// 受信ループ
	for {
		// チャネルからの受信
		// クローズされたか知る必要がなければ２番目の値を受け取る必要はない
		value, ok := <-c

		// チャネルがクローズされるとokにfalseが返される
		if !ok {
			fmt.Println("END")
			break
		}

		fmt.Println(value)
	}
}
