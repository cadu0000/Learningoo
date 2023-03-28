package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador = 0

const goRoutines = 100

func main() {

	criarRoutines(goRoutines)
	wg.Wait()
	fmt.Println(goRoutines, contador)
	// valores diferentes pois mais de uma thread do processador
	// está utilizando uma mesma variável ao mesmo tempo
}

func criarRoutines(i int) {
	wg.Add(i)
	for k := 0; k < i; k++ {
		go func() {
			j := contador
			runtime.Gosched() // habilita outra thread do processador
			j++
			contador = j
			wg.Done()
		}()
	}
}
