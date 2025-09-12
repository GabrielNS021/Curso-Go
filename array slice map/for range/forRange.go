package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador faz a contagem

	for i, numero := range numeros {
		fmt.Printf("%d: %d\n", i+1, numero)
	}

	// ignora o indice
	for _, num := range numeros { // se remover o _ ele pega a posicao do indice
		fmt.Println(num)
	}
}