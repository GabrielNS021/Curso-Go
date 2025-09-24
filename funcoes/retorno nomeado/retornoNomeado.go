package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // return limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	fmt.Println(trocar(5, 6))
}