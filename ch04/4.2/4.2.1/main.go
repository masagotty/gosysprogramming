package main

import (
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	//終了を受け取るためのチャネル
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		// バッファなしチャネルでは、送信時は受け取り側が受信操作(今回は"<-done")をするまで停止するので、<-done時に関数が最後まで実行される？
		done <- true
	}()
	// 終了を待つ
	<-done
	fmt.Println("all tasks are finished!")
}
