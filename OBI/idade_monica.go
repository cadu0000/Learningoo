package main

import "fmt"

func main() {
	var (
		m, a, b, c, maisVelho int
	)

	fmt.Scan(&m, &a, &b)

	c = -((a + b) - m)
	filhos := []int{a, b, c}

	for i := 0; i < len(filhos); i++ {
		if filhos[i] > maisVelho {
			maisVelho = filhos[i]
		}
	}
	fmt.Print(maisVelho)
}
