package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")
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

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
