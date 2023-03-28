package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falarNome() { //retorna o método falarNome() que é implementado pela interface humanos
	fmt.Println("o meu nome é:", p.nome)
}
func (p *pessoa) verificarIdade() { //retorna o método verificarIdade() que é implementado pela interface humanos
	if p.idade < 18 {
		fmt.Println("tenho", p.idade, "anos, portanto sou menor de idade, então não posso ir a festa")
	} else {
		fmt.Println("tenho", p.idade, "anos, então to liberado pra festejar")
	}
}

type humanos interface {
	falarNome()
	verificarIdade()
}

func nomeIdade(h humanos) {
	h.falarNome()
	h.verificarIdade()
}

func main() {
	pessoa1 := pessoa{"Lucas", 17}
	pessoa2 := pessoa{"Gabriel", 21}

	nomeIdade(&pessoa1)
	nomeIdade(&pessoa2)

	pessoa1 = pessoa{"Las", 15}
	nomeIdade(&pessoa1)
}
