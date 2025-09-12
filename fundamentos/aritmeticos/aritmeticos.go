package main

import (
	"fmt"
	m "math"	
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =", a+b)
	fmt.Println("Subtracao =", a-b)
	fmt.Println("Divisao =", a/b)
	fmt.Println("Multiplicacao =", a*b)
	fmt.Println("Modulo =", a%b)

	//bitwise
	fmt.Println("E =", a&b) // 11 & 10 = 10
	fmt.Println("Ou =", a|b) // 11 | 10 = 11
	fmt.Println("Xor =", a^b) // 11 ^10 = 01

	// math
	c := 3.0
	d := 2.0

	fmt.Println("Maior =", m.Max(float64(a), float64(b)))
	fmt.Println("Menor =", m.Min(c, d))
	fmt.Println("Exponencial =", m.Pow(c, d))
}