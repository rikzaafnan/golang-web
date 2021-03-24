package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	mux := http.NewServeMux()

	// abouthandler := func(w http.ResponseWriter, r *http.Request) {

	// 	w.Write([]byte("About Page"))

	// }

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)
	mux.HandleFunc("/product", productHandler)
	// mux.HandleFunc("/about", abouthandler)
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Profile"))
	// })

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

	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("welcome to root"))

}

func productHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)

		return
	}

	// w.Write([]byte("Product Page"))

	fmt.Fprintf(w, "Product Page : %d", idNumb)

}
