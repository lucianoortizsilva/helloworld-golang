package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	//concorrencia != paralelismo
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done()
	}()

	go func() {
		escrever("Acho que é isso")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
