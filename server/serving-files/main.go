package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", picture)
	http.HandleFunc("/serve", servePicture)
	http.ListenAndServe(":3001", nil)
}

//not serve
func picture(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="https://cdn-images-1.medium.com/max/1400/1*S7VjOUzpgIORXUIOa17UKg.jpeg">`)
}

func servePicture(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("go.png")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	io.Copy(w, f)
}
