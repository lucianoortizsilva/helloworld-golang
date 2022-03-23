package main

import "fmt"

func recuperarApicacao() {
	if r := recover(); r != nil { // RECOVER
		fmt.Println("RECUPERANDO APLICAÇÃO COM SUCESSO")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarApicacao()

	media := (n1 + n2) / 2

	if media >= 6 {
		fmt.Println("PARABÉNS! ALUNO APROVADO!!!")
		return true
	} else if media > 0 {
		fmt.Println("OPS! ALUNO REPROVADO!!!")
		return false
	}

	panic("ATENÇÃO!!! ALUNO MUITO BURRO.")
}

func main() {

	fmt.Println(alunoEstaAprovado(6, 6)) // TRUE
	fmt.Println(alunoEstaAprovado(6, 5)) // FALSE
	fmt.Println(alunoEstaAprovado(0, 0)) // PANIC -----+ RECOVER
	fmt.Println("PÓS EXECUTAR TUDO")
}
