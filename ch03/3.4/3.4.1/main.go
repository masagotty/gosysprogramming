package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5)
		// os.Stdin.Readで読み込んだ内容をbufferにぶっこむ
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}
