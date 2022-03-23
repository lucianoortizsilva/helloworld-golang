package main

import "fmt"

func funcao1() {
	fmt.Println("EXECUTANDO FUNÇÃO 1")
}

func funcao2() {
	fmt.Println("EXECUTANDO FUNÇÃO 2")
}

func isAlunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada com sucesso")
	fmt.Println("Verificando se aluno esta aprovado!")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer funcao1() // defer = adiar
	funcao2()
	isAlunoAprovado(4, 9)
}
