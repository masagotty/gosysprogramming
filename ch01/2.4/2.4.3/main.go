package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	// 連続でbuffer.Writeすればためておける
	// buffer.Write([]byte("aiuoe\n"))
	fmt.Println(buffer.String())
}
