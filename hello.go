package main

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"log"
	"net/http"
)

func GenerateQRcode(phrase string) {
	err := qrcode.WriteFile(phrase, qrcode.Medium, 256, "img/qr.png")
	if err!= nil {
		log.Fatal(err)
	}
}

func HandlerBill (w http.ResponseWriter, r *http.Request){
	URLparam := r.URL.Path[1:]
	if len(URLparam) != 0 {
				GenerateQRcode(URLparam)
				fmt.Println(URLparam)
				fmt.Fprintf(w, "<h1>Here is the QRcode for %s</h1> <a href='img/qr.png' download>Dowload it</a><img src='/img/qr.png'> ", URLparam )
			}else{
		fmt.Fprintf(w,"Hello bill !\nNothing to see here.\n")
	}
}

func HandlerFavicon (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello I am Favicon!")

}



func main() {
	fmt.Printf("Hello, world.\n")
	http.HandleFunc("/", HandlerBill)
	http.HandleFunc("/favicon.ico", HandlerFavicon)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe(":8080", nil)


}

