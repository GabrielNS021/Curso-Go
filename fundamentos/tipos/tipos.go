package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Inteiros
	fmt.Println("Tipo de 1000 é:", reflect.TypeOf(1000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
	
	i1 := math.MaxInt64
	fmt.Println("O valor maximo é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))


	var i2 rune = 'a' // representa um mapameamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("Tipo de x", reflect.TypeOf(x))
	fmt.Println("Tipo de x", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println(!bo)

	//String
	s1 := "Ola mundo"
	fmt.Println(s1 + "!")
	fmt.Println("Tamanho texto:", len(s1))
	
	// String multiplas linhas ``
	s2 := `Ola
Mundo`
	fmt.Println(s2)
}