package main

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"log"
	"net/http"
)

func GenerateQRcode(phrase string) {
	err := qrcode.WriteFile(phrase, qrcode.Medium, 256, "qr.png")
	if err!= nil {
		log.Fatal(err)
	}
}

func HandlerBill (w http.ResponseWriter, r *http.Request){
	URLparam := r.URL.Path[1:]
	if len(URLparam) != 0 {
		fmt.Fprintf(w, "Hello %s", URLparam )
		GenerateQRcode(URLparam)
		http.ServeFile(w, r, "qr.png")
	} else {
		fmt.Fprintf(w,"Hello bill !\nNothing to see here.\n")
	}	
}


func main() {
	fmt.Printf("Hello, world.\n")
	http.HandleFunc("/", HandlerBill)
	http.ListenAndServe(":8080", nil)

}

