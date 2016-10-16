package main

import "log"

func main() {
	sem := make(chan struct{}, 4) // セマフォ＝コンカレンシー＝並行度

	for i := 0; i < 42; i++ {
		sem <- struct{}{} // 空構造体で１枠確保

		go func(i int, s chan struct{}) {
			log.Println("goroutine: ", i)
			<-s // 確保していた１枠を解放
		}(i, sem)
	}

	log.Println("the main goroutine")
}
