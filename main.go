package main

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GenerateQRcode(phrase string) ([]byte) {
	png, err := qrcode.Encode(phrase, qrcode.Medium, 256)
	if err!= nil {
		log.Fatal(err)
	}
	return png
}

func HandlerBill (w http.ResponseWriter, r *http.Request){
	URLparam := r.URL.Path[1:]
	if len(URLparam) != 0 {
				fmt.Fprintf(w, "<h1>Here is the QRcode for %s</h1> <a href='/img/%s' download><button>Download!</button></a> <img src='/img/%s'> ", URLparam, URLparam, URLparam )
			}else{
		fmt.Fprintf(w,"Hello bill !\nNothing to see here.\n")
	}
}

func HandlerFavicon (w http.ResponseWriter, r *http.Request) {
	fmt.Print("demande de favicon")
	http.Redirect(w, r, "/img/favicon.ico", http.StatusFound)

}

func Handlerimage (w http.ResponseWriter, r *http.Request) {
	URLparam := (r.URL.Path[5:])
	writeImage(w, GenerateQRcode(URLparam))
}

func writeImage(w http.ResponseWriter, png []byte) {

	w.Header().Set("Content-Type", "image/png")
	if _, err := w.Write(png); err != nil {
		log.Println("unable to write image.")
	}
}



func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	//http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/img/", Handlerimage)
	http.HandleFunc("/", HandlerBill)
	http.HandleFunc("/favicon.ico", HandlerFavicon)
	http.ListenAndServe(":"+ port, nil)

}

