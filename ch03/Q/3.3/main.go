package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	fileA, err := zipWriter.Create("testA.txt")
	if err != nil {
		panic(err)
	}
	textA := strings.NewReader("1つ目やで")
	io.Copy(fileA, textA)

	fileB, err := zipWriter.Create("testB.txt")
	if err != nil {
		panic(err)
	}
	textB := strings.NewReader("2つ目やで")
	io.Copy(fileB, textB)
}
