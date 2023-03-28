package main

import "fmt"

const total = 20

func main() {
	impar := make(chan int)
	par := make(chan int)
	verificador := make(chan bool)

	go filtroParidade(par, impar, verificador)
	verificar(par, impar, verificador)
}

func filtroParidade(par, impar chan int, verificador chan bool) {
	for i := 0; i < total; i++ {
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
	verificador <- true
}
func verificar(par, impar chan int, verificador chan bool) {
	for {
		select {
		case valor := <-par:
			fmt.Println("o número", valor, "é par")
		case valor := <-impar:
			fmt.Println("o número", valor, "é ímpar")
		case valor, ok := <-verificador:
			if !ok { // comma ok para evitar o envio de valores inexistentes
				fmt.Println("erro", valor)
			} else {
				fmt.Println("finalizando?", valor)
				return
			}

		}

	}
}
