package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {

	a := 1
	b := 2

	imprimir((somar(a, b)))
}