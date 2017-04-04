package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"log"
	"net/http"
)


func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := version.PatternSize()
	img := image.NewNRGBA(image.Rect(0,0,size,size))
	return png.Encode(w,img)
	}

type Version int8
func (v Version) PatternSize() int {
 	return 4*int(v) + 17
	}

func main() {
	fmt.Printf("Hello, world.\n")
	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = GenerateQRCode(file, "555-2368", 1)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world",)
	})
	http.ListenAndServe(":8080", nil)
}
