package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("タイマー開始！")
	t := time.After(1 * time.Second)
	<-t
	fmt.Println("1秒経ったお")
}
