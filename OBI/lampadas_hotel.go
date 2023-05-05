package main

import "fmt"

func main() {
	var (
		iA, iB, fA, fB bool
	)
	var contador int
	fmt.Scan(&iA, &iB, &fA, &fB)

	if iA != fA && iB != fB {
		contador++
	} else {
		contador = 2
	}
	fmt.Print(contador)
}
