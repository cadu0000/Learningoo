package main

import (
	"fmt"
	"sync"
)

const total = 20

// pesquisar sobre bubble sort
var pares = []int{}
var impares = []int{}

var wg sync.WaitGroup

//var mu sync.Mutex

func main() {
	testarParidade(total)
	wg.Wait()
}

func testarParidade(x int) {

	for i := 0; i < x; i++ {
		wg.Add(1)
		numero := i
		go func() {
			switch {
			case numero == 0:
				pares = append(pares, numero)
				fmt.Printf("número %v é nulo\n", numero)
			case numero%2 == 0:
				impares = append(impares, numero)
				fmt.Printf("número %v é par\n", numero)
			default:
				fmt.Printf("número %v é ímpar\n", numero)
			}
			wg.Done()
		}()
	}
}
