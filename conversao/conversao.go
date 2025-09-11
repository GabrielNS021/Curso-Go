package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	// arredondar
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	// fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste "+ strconv.Itoa(123))
	
	// String para Int
	num, _ := strconv.Atoi("123") // Primeiro eh o valor e o segundo eh o erro
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}
}