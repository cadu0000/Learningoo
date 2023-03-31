package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador = 0
var mu sync.Mutex

const goRoutines = 100

func main() {

	criarRoutines(goRoutines)
	wg.Wait()
	fmt.Println(goRoutines, contador)
	// valores iguais pois o tipo mutex garante os médotodos Lock() e Unlock ()
	// que garantem que uma variável não possa ser manipulada por mais de uma thread ao mesmo tempo
}

func criarRoutines(i int) {
	wg.Add(i)
	for k := 0; k < i; k++ {
		go func() {
			mu.Lock() // bloqueia a ação de outras threads
			j := contador
			runtime.Gosched() // habilita outra thread do processador
			j++
			contador = j
			mu.Unlock() // desbloqueia a ação de outras threads
			wg.Done()
		}()
	}
}
