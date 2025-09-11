package main

import "fmt"
import m "math"

func main() {

	/* const/var (
		a = 1
		b = 2
	) */

	// g, h, i := 2, false, "epa!"


	const PI float64 = 3.1415
	var raio = 3.2
	area := PI * m.Pow(raio, 2)

	fmt.Println(area)
}