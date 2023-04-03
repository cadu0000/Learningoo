package main

import (
	"encoding/json"
	"fmt"
)

type contato struct {
	Nome   string
	Numero int
}

func main() {
	contato1 := contato{"Jõao", 99852426}
	contato2 := contato{"Lucas", 99245786}
	contato3 := contato{"gabi", 99578420}

	listaContatos := []contato{contato1, contato2, contato3}
	fmt.Println(listaContatos)

	sb, err := json.Marshal(listaContatos)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(sb))

	contatos := []contato{}

	err = json.Unmarshal(sb, &contatos)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(contatos)
}
