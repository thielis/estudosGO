package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Arrays tem tamanho fixo
	var notas [3]float64              // declaração do array
	fmt.Println(notas)                // print do array zerado
	notas = [3]float64{1.2, 3.4, 5.6} //atribuição
	fmt.Println(notas)
	fmt.Println("Apresentando tipo de variável:", reflect.TypeOf(notas))

	//Exemplo atribuição de cada elemento de forma individual
	var notasInt [3]int

	notasInt[0] = 1
	notasInt[1] = 2
	notasInt[2] = 3
	fmt.Println(notasInt)
	//OU
	notasInt[0], notasInt[1], notasInt[2] = 11, 12, 13
	fmt.Println(notasInt)

	//Declaração curta definindo o tamanho do array
	letras := [3]string{"a", "b", "c"} // esta é uma forma
	fmt.Println(letras)

	//Nesta forma o tamanho é definido pelo compilador
	//O compilador conta a quantidade de elementos e define o tamanho do array
	//A notação "[...]" que identifica esse modo de funcionamento
	numeros := [...]int{10, 20, 30, 40, 50}
	fmt.Println(numeros)

}
