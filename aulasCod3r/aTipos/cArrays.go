package main

import "fmt"

func main() {
	//Arrays tem tamanho fixo
	var notas [3]float64              // declaração do array
	fmt.Println(notas)                // print do array zerado
	notas = [3]float64{1.2, 3.4, 5.6} //atribuição
	fmt.Println(notas)

	var notasInt [3]int
	//atribuição de cada elemento de forma individual
	notasInt[0] = 1
	notasInt[1] = 2
	notasInt[2] = 3
	fmt.Println(notasInt)
	//OU
	notasInt[0], notasInt[1], notasInt[2] = 11, 12, 13
	fmt.Println(notasInt)

	//declaração curta
	letras := [3]string{"a", "b", "c"}
	fmt.Println(letras)

	//Slices ou partes de um array
	//não tem limite de itens

}
