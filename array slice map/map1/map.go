package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)
	
	aprovados[123] = "Maria"
	aprovados[132] = "Pedro"
	aprovados[321] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123])
	delete(aprovados, 123)
	fmt.Println(aprovados[123])
}