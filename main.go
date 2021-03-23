package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	log.Println("Starting web on port 80808")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world, saya sedang belajr golang Web"))

}
