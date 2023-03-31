package main

import (
	"fmt"
	"sync"
)

const total = 20

func main() {
	impar := make(chan int)
	par := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	
	go filtroParidade(par, impar, &wg)
	exibirResultado(par, impar)
	wg.Wait()

}

func filtroParidade(par, impar chan int, wg *sync.WaitGroup) { //}, verificador chan bool) {
	for i := 1; i < total; i++ {
		if i == 0 {
			fmt.Println("O número", i, "é nulo")
		} else if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	wg.Done()
}

func exibirResultado(par, impar chan int) {
	for {
		select {
		case valor := <-par:
			if valor == 0 {
				return
			}
			fmt.Println("o número", valor, "é par")
		case valor := <-impar:
			if valor == 0 {
				return
			}
			fmt.Println("o número", valor, "é ímpar")
		}
	}
}
