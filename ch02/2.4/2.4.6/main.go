package main

// io.MultiWriter
// import (
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("multiwriter.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	writer := io.MultiWriter(file, os.Stdout)
// 	io.WriteString(writer, "io.MultiWriter example\n")
// }

// 圧縮
// import (
// 	"compress/gzip"
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("test.txt.gzip")
// 	if err != nil {
// 		panic(err)
// 	}
// 	writer := gzip.NewWriter(file)
// 	writer.Header.Name = "test.txt"
// 	io.WriteString(writer, "gzip.Writer example\n")
// 	writer.Close()
// }

// bufio.Writer
import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer\n")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
