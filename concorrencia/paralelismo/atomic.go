package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

const goRoutines = 1000

func main() {

	criarRoutines(goRoutines)
	wg.Wait()
	fmt.Println(goRoutines, contador)
}

func criarRoutines(i int) {
	wg.Add(i)
	for k := 0; k < i; k++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			wg.Done()
		}()
	}
}
