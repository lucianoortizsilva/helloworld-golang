package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	} else {
		retorno := fibonacci(posicao-2) + fibonacci(posicao-1)
		return retorno
	}
}

func main() {

	//1 1 2 3 5 8 13 ...
	valorFibonacci := uint(12)

	for i := uint(1); i <= valorFibonacci; i++ {
		fmt.Println(fibonacci(i))
	}

}
