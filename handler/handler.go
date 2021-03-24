package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {

		log.Println(err)

		http.Error(w, "Error is happening, keep calm ", http.StatusInternalServerError)
		return

	}

	data := map[string]interface{}{
		"title":   "I'm Learning Golang Web",
		"content": "I'm Learning Golang web with Agung Setiawan",
	}

	err = tmpl.Execute(w, data)
	if err != nil {

		log.Println(err)

		http.Error(w, "Error is happening, keep calm ", http.StatusInternalServerError)
		return

	}

	// w.Write([]byte("welcome to root"))

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world, saya sedang belajr golang Web"))

}

func MarioHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("in iroute dari mario"))

}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)

		return
	}

	// w.Write([]byte("Product Page"))

	fmt.Fprintf(w, "Product Page : %d", idNumb)

}
