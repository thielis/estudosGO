package main

import (
	"fmt"
	"time"
)

//For é utilizado para laços

func lacoInfinito() {
	// Laço infinito
	num := 0 //É preciso declarar a variável antes do for
	for {
		fmt.Println("Laço infinito contagem:", num)
		num++
		time.Sleep(time.Second)
	}
}

func while(i, j int) {
	//Go não tem While então utilizamos o For para isso
	//i define quantos loops
	for i <= j {
		fmt.Println(i)
		i++
		// time.Sleep(time.Second) //atras em 1s cada iteração
	}
}

func fornormal(iteracoes, incremento int) {
	//For normal
	for j := 0; j <= iteracoes; j += incremento {
		fmt.Println(j)
		// time.Sleep(time.Second)
	}
}

// Equivale ao foreach de outras linguagens
func foreach(numeros []int) { //pode-se definir outro tipo de dado
	for i, num := range numeros {
		//i é variável de iteração ou índice, se não for utilizar é preciso suprimir com "_"
		//num é a variável que recebe o valor da posição do array em cada ciclo
		//numeros é o array que será consultado em que utilizando range ele percorre a cada índice
		fmt.Printf("Índice %d): %d \n", i, num)
	}
}

// Também é possível utilizar o for suprimindo a variável do índice
func soletrando(palavra string) { //for sempre retorna o índice
	for _, letra := range palavra { //Suprimindo a variável do índice para não precisar utilizá-la
		fmt.Println(string(letra)) // neste caso é necessário converter a variável pois vai retornar o número ref. tabela ASCII
	}
}

func soletrando2(palavra string) {
	for letra := range palavra { //Suprimindo a variável do índice para não precisar utilizá-la
		fmt.Println(letra) //Neste caso o for retornará através da variável o índice e não o valor contido,
		//ou seja, não adianta fazer a conversão como na função soletrando().
	}
}

func main() {

	fmt.Println("For iterando de 1 em 1")
	fornormal(10, 1)
	fmt.Println("----------------------")

	fmt.Println("For iterando de 1 em 1")
	fornormal(10, 2)
	fmt.Println("----------------------")

	fmt.Println("Enquando(While) menor que 10")
	while(0, 10)
	fmt.Println("----------------------")

	numarray := []int{1, 3, 5, 6}
	foreach(numarray)
	fmt.Println("----------------------")

	soletrando("paralelepipedo")
	fmt.Println("----------------------")

	soletrando2("paralelepipedo")
	fmt.Println("----------------------")

	// For infinito
	lacoInfinito()

}
