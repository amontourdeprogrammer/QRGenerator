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

func HandlerBill (w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Hello bill :", r)
	GenerateQRcode(r.URL.Path)

}

func main() {
	fmt.Printf("Hello, world.\n")
	http.HandleFunc("/", HandlerBill)
	http.ListenAndServe(":8080", nil)
}

