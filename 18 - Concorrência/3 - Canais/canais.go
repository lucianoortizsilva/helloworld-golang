package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i >= 3 {
			canal <- texto
			time.Sleep(time.Second)
		}
	}
	close(canal)
}

func main() {
	canal := make(chan string)
	go escrever("Olá mundo!!!", canal)

	fmt.Println("depois da func escrever")

	for {
		mensagem, aberto := <-canal
		if aberto {
			fmt.Println(mensagem)
		} else {
			break
		}
	}

	fmt.Println("ACABOOOOOOOOOOOOOOOOOU")
}
