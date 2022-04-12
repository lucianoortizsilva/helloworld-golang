package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func init() {
	config.Carregar()
	cookies.Configurar()
}

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()
	fmt.Printf("Rodando WebApp na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
