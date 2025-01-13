package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Teste")
	a, b := 2, 3
	x := soma(a, b)
	fmt.Printf("A soma de %d e %d é igual a: %d", a, b, x)

}
