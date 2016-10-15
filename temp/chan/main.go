package main

import "fmt"

func main() {
	// ゴルーチン間でデータの送受信に使うチャネルを生成
	ch := make(chan int)

	// ゴルーチン（＝別スレッド）
	// sという変数名でチャネルに送信する [ chan<- 型 ]=チャネルに対して指定の型で送信するという意味
	go func(s chan<- int) {
		for i := 0; i < 5; i++ {
			s <- i // 0～4までの値を順次、s という変数名を通してチャネルに送信！
		}
		// チャネルは不要になったタイミングでクローズ！
		close(s)
	}(ch) // 実際に送信するチャネルは冒頭で生成した ch

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
}
