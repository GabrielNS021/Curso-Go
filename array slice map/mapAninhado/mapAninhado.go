package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"Gabriela Silva": 1234.60,
			"Guga Pereira": 4321.50,
		},
		"J": {
			"Jose Joao": 9876.30,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
