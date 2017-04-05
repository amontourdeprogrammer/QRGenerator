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
		if URLparam == "favicon.ico"{
		fmt.Fprintf(w, "<html><img src='http://www.artsmartiauxcombat.com/images/favicon.ico-16x16-i11.ico'></html>")
		}else{
			if URLparam == "img/qr.png"{
				fmt.Fprintf(w, "<html><img src='http://www.artsmartiauxcombat.com/images/favicon.ico-16x16-i11.ico'></html>")
			}else{
				GenerateQRcode(URLparam)
				fmt.Println(URLparam)
				fmt.Fprintf(w, "<h1>Hello</h1> <img src='http://localhost:8080/img/qr.png'> %s", URLparam )
			}
	} else {
		fmt.Fprintf(w,"Hello bill !\nNothing to see here.\n")
	}
}



func main() {
	fmt.Printf("Hello, world.\n")
	http.HandleFunc("/", HandlerBill)
	http.ListenAndServe(":8080", nil)

}

