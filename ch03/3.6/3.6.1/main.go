package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目`

// ながいやつ
// func main() {
// 	reader := bufio.NewReader(strings.NewReader(source))
// 	for {
// 		line, err := reader.ReadString('\n')
// 		fmt.Printf("%#v\n", line)
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// 短いやつ
func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
