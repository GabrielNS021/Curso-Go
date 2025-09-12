package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3} // slice

	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// a2 := [5]int{1, 2, 3, 4, 5}


}