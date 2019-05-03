package main

import (
	"encoding/csv"
	"os"
)

func main() {
	data := []string{"hoge", "fuga", "foo", "bar"}

	writer := csv.NewWriter(os.Stdout)
	writer.Write(data)
	writer.Flush()
}
