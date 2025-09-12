package main

import "fmt"

// Go nao tem aritmetica de ponteiro

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereco da variavel
	*p++
	
	fmt.Println("I:", i)
	fmt.Println("P:", p) // Endereco
	fmt.Println("P:", *p) // Valor
}