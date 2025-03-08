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
		time.Sleep(time.Second) //atras em 1s cada iteração
	}
}

func fornormal(iteracoes, incremento int) {
	//For normal
	for j := 0; j <= iteracoes; j += incremento {
		fmt.Println(j)
		time.Sleep(time.Second)
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
	// For infinito
	lacoInfinito()

}
