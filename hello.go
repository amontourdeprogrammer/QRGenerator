package main

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"log"
	"net/http"
)

func GenerateQRcode(phrase string) {
	err := qrcode.WriteFile(phrase, qrcode.Medium, 256, "qr.png")
	fmt.Printf("err=", err)

	if err!= nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Printf("Hello, world.\n")
	GenerateQRcode("je suis Sarah")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world",)
	})
	http.ListenAndServe(":8080", nil)
}

