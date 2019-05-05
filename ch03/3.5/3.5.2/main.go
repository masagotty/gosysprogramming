package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32ビットのビッグエンディアン：10000
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// エンディアンの変換
	//binary.Readで、bytes.NewReader(data)で読み込んだバイナリを変換しiにぶっこむ
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
