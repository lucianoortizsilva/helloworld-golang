package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {

	var u usuario
	u.nome = "Luciano Ortiz"
	u.idade = 38
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua qualquer", 123}

	outroUsuario := usuario{"Mariana", 10, enderecoExemplo}
	fmt.Println(outroUsuario)

	maisUmUsuario := usuario{nome: "Usuario sem nome"}
	fmt.Println(maisUmUsuario)

}
