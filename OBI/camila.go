package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		a, b, c int
	)
	fmt.Scan(&a, &b, &c)
	idades := []int{a, b, c}
	sort.Ints(idades)
	camila := idades[1]
	fmt.Print(camila)

}
