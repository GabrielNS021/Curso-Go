package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("Parametro 1: %v\nParametro 2: %v\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %v, %v", p1, p2)
}

func f5() (string, string) {
	return "ABC", "DEF"
}

func main() {
	f1()
	f2("Ola", "Mundo")
	
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Printf("%v, %v\n", r3, r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}