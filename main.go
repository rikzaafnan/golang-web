package main

import (
	"golang-web/handler"
	"log"
	"net/http"
	"os"
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
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile"))
	})
	port := os.Getenv("PORT")

	if port == "" {
		port = "8181"
	}

	log.Println("server run port :", port)

	err := http.ListenAndServe(":"+port, mux)

	log.Fatal(err)

}
