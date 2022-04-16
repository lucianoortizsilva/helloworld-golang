package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	fmt.Println(">>> REX: ")
	cachorroRex := cachorro{"Rex", "Vira-lata", 12}
	cachorroRexJSON, erro := json.Marshal(cachorroRex)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroRexJSON)
	fmt.Println(bytes.NewBuffer(cachorroRexJSON))
	///////////////////////////

	fmt.Println(">>> TOBI: ")
	cachorroTobi := map[string]string{
		"nome": "Tobi",
		"raca": "pooodle",
	}
	cachorroTobiJSON, erro := json.Marshal(cachorroTobi)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroTobiJSON)
	fmt.Println(bytes.NewBuffer(cachorroTobiJSON))
}
