package main

import (
	"fmt"
	"time"
)

func main() {
	c := 0

	for c  <= 10 {
		fmt.Println(c)
		c++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		// Laco infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second*5)
	}
}