package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// abouthandler := func(w http.ResponseWriter, r *http.Request) {

	// 	w.Write([]byte("About Page"))

	// }
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile"))
	})
	// mux.HandleFunc("/", handler.HomeHandler)
	// mux.HandleFunc("/hello", handler.HelloHandler)
	// mux.HandleFunc("/mario", handler.MarioHandler)
	// mux.HandleFunc("/product", handler.ProductHandler)
	// mux.HandleFunc("/about", abouthandler)
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Profile"))
	// })

	log.Println("server run port :8181")

	err := http.ListenAndServe(":8181", mux)

	log.Fatal(err)

}
