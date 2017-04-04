package main

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Hello, world.\n")
	var png []byte
	png, err := qrcode.Encode("Hello", qrcode.Medium, 256)
	err = qrcode.WriteFile("Hello", qrcode.Medium, 256, "qr.png")
	if err!= nil {
		log.Fatal(err)
	}
	_ = png

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world",)
	})
	http.ListenAndServe(":8080", nil)
}
