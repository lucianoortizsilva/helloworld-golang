package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {

	//chave := make([]byte, 64)

	//if _, erro := rand.Read(chave); erro != nil {
	//log.Fatal(erro)
	//}
	//txtBase64 := base64.StdEncoding.EncodeToString(chave)
	//fmt.Println(txtBase64)
}

func main() {

	config.Carregar()
	//fmt.Println(config.SecretKey)
	fmt.Printf("Rodando Api na porta %d\n", config.Porta)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
