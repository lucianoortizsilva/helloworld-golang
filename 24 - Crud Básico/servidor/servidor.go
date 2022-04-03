package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}
	var u usuario

	if erro = json.Unmarshal(body, &u); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	fmt.Println(body)
}
