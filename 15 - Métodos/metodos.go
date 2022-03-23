package main

import (
	"fmt"
)

type usuario struct {
	nome    string
	apelido string
	idade   uint8
}

func (user usuario) salvarDadosDoUsuario() {
	fmt.Printf("Salvando dados do usuário %s na base de dados de produção\n", user.nome)
}

func (user *usuario) alterarApelidoDoUsuario() {
	user.apelido = "Kbça branca"
}

func (user *usuario) usuarioMaiorDeIdade() bool {
	if user.idade >= 18 {
		return true
	} else {
		return false
	}
}

func main() {

	user1 := usuario{"Luciano", "Kbça de coelho", 38}
	user1.salvarDadosDoUsuario()
	user1MaiorDeIdade := user1.usuarioMaiorDeIdade()
	fmt.Println(user1)
	fmt.Printf("Usuário 1 é maior de idade ? ")
	fmt.Println(user1MaiorDeIdade)
	fmt.Println("------------")

	user2 := usuario{"Dalessandro", "Galo cinza", 17}
	user2MaiorDeIdade := user2.usuarioMaiorDeIdade()
	fmt.Println(user2)
	user2.alterarApelidoDoUsuario()
	fmt.Println(user2)
	fmt.Printf("Usuário 2 é maior de idade ? ")
	fmt.Println(user2MaiorDeIdade)

}
