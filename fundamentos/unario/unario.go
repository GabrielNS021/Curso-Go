package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x ++ // +1
	y --
	fmt.Println(x, y)

	// ERRADO: fmt.Println(x == y--)
}