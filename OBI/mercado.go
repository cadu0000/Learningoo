package main

import "fmt"

func main() {

	var (
		p, g, proportion float64
	)
	var n int
	var menorPorcentagem = 1000.00
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		fmt.Scan(&g)

		proportion = p / g

		if proportion <= menorPorcentagem {
			menorPorcentagem = proportion
		}
	}
	fmt.Print(menorPorcentagem * 1000)
}
