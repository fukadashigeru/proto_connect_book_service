// server/main.go
package main

import (
	"log"
	"net/http"

	"proto_connect_book_service/bookapi" // この行を追加

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	bookapi.HandleBook(mux) // bookapi パッケージの関数を使用

	log.Fatal(http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	))
}
