package main

import "fmt"

func main() {
	//Operadores Unários no Go só tem posfixado
	//Atribui valor após a leitura da veriável
	x, y := 1, 2

	//Operador Unário de incremento
	x++
	fmt.Println("O valor atual de 'X' é: ", x)
	//Operador Unário de decremento
	y--
	fmt.Println("O valor atual de 'X' é: ", y)

}
