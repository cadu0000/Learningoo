package main

import "fmt"

func main() {
	nomeIdade := map[string]int{
		"Jonas":     15,
		"Guilherme": 17,
		"Letícia":   22,
		"Clara":     18,
		"Cadu":      18,
	}
	verificarIdade(nomeIdade)
}

func verificarIdade(x map[string]int) {
	for k, v := range x { //os valores vêm em ordem aleatória
		switch {
		case v >= 18:
			fmt.Printf("%v pode entrar na festa porque tem %v anos\n", k, v)
		case v < 18:
			fmt.Printf("%v não pode entrar na festa porque tem %v anos\n", k, v)
		}
	}
}
