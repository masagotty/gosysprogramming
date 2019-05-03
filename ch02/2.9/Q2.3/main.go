package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	///json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	//ここにコードを書く
	compressor := gzip.NewWriter(w)
	writer := io.MultiWriter(compressor, os.Stdout)
	// gzip.NewWriter()とos.StdoutにJSONを書き込むために、json.NewEncoder()を使用
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	// encoderにsourceをぶっこむとwriterに出力される
	encoder.Encode(source)
	compressor.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
