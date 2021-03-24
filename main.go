package main

import (
	"golang-web/handler"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// abouthandler := func(w http.ResponseWriter, r *http.Request) {

	// 	w.Write([]byte("About Page"))

	// }

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	// mux.HandleFunc("/about", abouthandler)
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Profile"))
	// })

	log.Println("Starting web on port 80808")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
