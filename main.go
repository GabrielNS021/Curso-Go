package main

import "fmt"
import m "math"

func main() {

	const PI float64 = 3.1415
	var raio = 3.2
	area := PI * m.Pow(raio, 2)

	fmt.Println(area)
}