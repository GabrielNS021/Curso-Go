package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64 {
		"Jose Joao": 11325.25,
		"Gabriela Silva": 15345.23,
		"Pedro Junior": 1200.0,
	}

	funcsESalarios["Rafael filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
