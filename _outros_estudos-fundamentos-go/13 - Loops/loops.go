package main

import (
	"fmt"
)

func main() {

	numero := 1

	for numero <= 10 {
		fmt.Println(numero)
		numero++
		//time.Sleep(time.Second)
	}

	for j := 50; j < 60; j += 2 {
		fmt.Println(j)
	}

	times := [5]string{"internacional", "vasco", "fluminense", "são paulo", "goiás"}

	for indice := 0; indice < len(times); indice++ {
		fmt.Println(times[indice])
	}

	fmt.Println("-------------------")

	for i, time := range times {
		fmt.Println(i, time)
	}

	fmt.Println("-------------------")

	for ind, animal := range "macaco" {
		fmt.Println(ind, string(animal))
	}

	//range na aceita struct
}
