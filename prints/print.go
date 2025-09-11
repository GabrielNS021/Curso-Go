package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.\n")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516
	xs := fmt.Sprint(x)

	// fmt.println("O valor de x é" + x) -> String != num
	fmt.Println("O valor de x é " + xs)

	fmt.Printf("O valor de x é %.2f\n", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("Inteiro: %d \nFloat: %f \nFloat 1 casa decimal: %.1f \nBool: %t \nString: %s\n", a, b, b, c, d)
	// %v = modo generico
	fmt.Printf("%v %v %v %v", a, b, c, d)

}