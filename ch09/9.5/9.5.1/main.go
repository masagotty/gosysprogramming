package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Temp file path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
}
