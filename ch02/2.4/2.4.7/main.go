package main

//Fprintf
// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
// }

//JSON
// import (
// 	"encoding/json"
// 	"os"
// )

// func main() {
// 	encoder := json.NewEncoder(os.Stdout)
// 	encoder.SetIndent("", "    ")
// 	encoder.Encode(map[string]string{
// 		"example": "encoding/json",
// 		"hello":   "world",
// 	})
// }

//net/http
import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
