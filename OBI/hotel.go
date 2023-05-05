package main

import "fmt"

func main() {
	var (
		d, a, n, total, estadia, aux int
	)

	fmt.Scan(&d, &a, &n)
	estadia = 32 - n
	if n > 15 {
		aux = 14
	} else {
		aux = n - 1
	}
	total = estadia * (d + aux*a)
	fmt.Print(total)
}
