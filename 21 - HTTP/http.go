package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregando usuários!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
