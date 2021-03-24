package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)

	log.Println("Starting web on port 80808")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world, saya sedang belajr golang Web"))

}

func marioHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("in iroute dari mario"))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("welcome to root"))

}
