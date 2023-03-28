package main

import (
	"fmt"
	"sort"
)

type carros struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carros

func (o ordenarPorPotencia) Len() int { return len(o) }
func (o ordenarPorPotencia) Swap(maiorPotencia, menorPotencia int) {
	o[maiorPotencia], o[menorPotencia] = o[menorPotencia], o[maiorPotencia]
}
func (o ordenarPorPotencia) Less(maiorPotencia, menorPotencia int) bool {
	return o[menorPotencia].potencia < o[maiorPotencia].potencia
	// esses três métodos são métodos da interface interface
}

type ordenarPorConsumo []carros

func (o ordenarPorConsumo) Len() int { return len(o) }
func (o ordenarPorConsumo) Swap(maiorConsumo, menorConsumo int) {
	o[maiorConsumo], o[menorConsumo] = o[menorConsumo], o[maiorConsumo]
}
func (o ordenarPorConsumo) Less(maiorConsumo, menorConsumo int) bool {
	return o[menorConsumo].consumo < o[maiorConsumo].consumo
}

func main() {
	meusCarros := []carros{
		{"palio", 70, 13},
		{"corsa", 72, 14},
		{"civic", 150, 9},
		{"lamborghini", 320, 3},
	}
	fmt.Println(meusCarros)
	sort.Sort(ordenarPorPotencia(meusCarros))
	fmt.Println(meusCarros)
	sort.Sort(ordenarPorConsumo(meusCarros))
	fmt.Println(meusCarros)

}
