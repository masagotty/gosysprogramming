package main

import (
	"crypto/rand"
	"io"
	"os"
)

// 俺の回答
// func main() {
// 	file, err := os.Create("rand.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	randomReader := rand
// 	lReader := io.LimitReader(randomReader, 1024)
// 	fmt.Println(lReader)
// }

// 回答
func main() {
	file, err := os.Create("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := io.LimitReader(rand.Reader, 1024)
	io.Copy(file, r)
}
