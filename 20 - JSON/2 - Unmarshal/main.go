package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"-"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//COVERTER PARA struct cachorro
	cachorroEmJSON := `{"nome":"Rex","raca":"Vira-lata","idade":12}`
	var c1 cachorro
	erro := json.Unmarshal([]byte(cachorroEmJSON), &c1)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c1)

	cachorro2EmJSON := `{"nome":"Rex2","raca":"Vira-lata KKK"}`
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}
