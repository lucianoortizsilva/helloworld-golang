package main

import "fmt"

type cachorro struct {
	nome string
	raca string
}

type gato struct {
	nome string
	raca string
}

type falando interface {
	falar() string
}

func (c cachorro) falar() string {
	if c.raca == "vira-lata" {
		return "AU AU AU AUUUUU"
	} else {
		return "AU"
	}
}

func (g gato) falar() string {
	if g.raca == "persa" {
		return "MIAU MIAU MIAUUUUUU"
	} else {
		return "MIAU"
	}
}

func emitirSom(f falando) {
	fmt.Printf("O animal falou %s", f.falar())
}

func main() {

	gatoPersa := gato{"Gatito", "persa"}
	emitirSom(gatoPersa)

	println("\n-------")

	cachorroLoko := cachorro{"Tobi", "vira-lata"}
	emitirSom(cachorroLoko)

}
