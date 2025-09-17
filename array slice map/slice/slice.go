package main

import (
	"fmt"
	"reflect"
)

// Slice nao eh um array
// Slice define um pedaco de array

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3} // slice

	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}


	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println(s3)

	// imaginar o slice como: tamanho e ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}