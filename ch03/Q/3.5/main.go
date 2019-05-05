package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("I'm masagotty")
	io.CopyN(os.Stdout, reader, 10)
}
