package main

import (
	"log"
	"net/http"
	"servidor"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandlerFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5000", router))

}
