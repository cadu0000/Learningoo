package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	idade     int
	orçamento float32
}

func (p *pessoa) nomePessoa() {
	fmt.Print("\nDigite seu nome:")
	_, err := fmt.Scan(&p.nome)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (p *pessoa) idadePessoa() {
	fmt.Print("Digite sua idade:")
	_, err := fmt.Scan(&p.idade)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (p *pessoa) orçamentoPessoa() {
	fmt.Print("Digite seu orçamento:")
	_, err := fmt.Scan(&p.orçamento)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (p *pessoa) verificadorDaBalada() {
	switch {
	case p.idade < 18 && p.orçamento > 100:
		fmt.Println(p.nome, "naõ está autorizado pois possui", p.idade, "anos, ou seja, é menor de idade")
	case p.idade >= 18 && p.orçamento > 100:
		fmt.Println(p.nome, " está autorizado pois possui ", p.idade, " anos, ou seja, é maior de idade e possui dinheiro pra entrada")
	case p.idade >= 18 && p.orçamento < 100:
		fmt.Println(p.nome, "não está autorizado pois possui", p.idade, " anos, ou seja, é maior de idade, porém não possui dinheiro pra entrada")
	case p.idade < 18 && p.orçamento < 100:
		fmt.Println(p.nome, "não está autorizado pois possui", p.idade, " anos, ou seja, é menor de idade e sequer possui dinheiro pra entrada")
	}
}

type balada interface {
	nomePessoa()
	idadePessoa()
	orçamentoPessoa()
	verificadorDaBalada()
}

func verificarDados(b balada) {
	b.nomePessoa()
	b.idadePessoa()
	b.orçamentoPessoa()
	fmt.Println("")
}
func entrarNaBalada(b balada) {
	b.verificadorDaBalada()
	fmt.Println("")
}
func checarNumeroDePessoas() []pessoa {
	arrayDePessoas := []pessoa{}
	numeroDePessoas := 0
	fmt.Print("Digite a quantidade de pessoas que desejam entrar:")
	_, err := fmt.Scan(&numeroDePessoas)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	for i := 0; i < numeroDePessoas; i++ {
		pessoas := pessoa{}
		verificarDados(&pessoas)
		arrayDePessoas = append(arrayDePessoas, pessoas)
	}
	for d := 0; d < len(arrayDePessoas); d++ {
		entrarNaBalada(&arrayDePessoas[d])
	}
	//fmt.Print(arrayDePessoas)
	return arrayDePessoas
}
func main() {
	fmt.Println("Bem-vindo ao verificador da balada")
	checarNumeroDePessoas()
}
