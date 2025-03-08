package main

import "fmt"

func soma(a, b int) int {
	//Função que retorna uma soma de duas entradas
	return a + b
}

func main() {
	fmt.Println("Teste")
	a, b := 2, 3
	x := soma(a, b)
	fmt.Printf("A soma de %d e %d é igual a: %d\n", a, b, x)
	fmt.Println("A soma de 2 e 3 é igual a:", soma(2, 3))

}
