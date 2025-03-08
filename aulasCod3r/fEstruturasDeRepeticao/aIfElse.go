package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Função randômica de valores para utilizar de testes
func numAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

// if, else if e else
func resultadoAluno(nota float64) string {
	if nota <= 3.0 && nota >= 0 {
		return "reprovado"
	} else if nota > 3.0 && nota <= 4.5 {
		return "recuperação"
	} else if nota > 4.5 {
		return "Aprovado"
	} else {
		return "valor invalido"
	}

}

//If com inicialização de uma variável

func interruptorLampada() {
	if i := numAleatorio(); i > 5 {
		fmt.Println("Luz Acesa")
	} else {
		fmt.Println("Luz Apagada")
	}
}

func main() {

	fmt.Println(resultadoAluno(-0.5))
	fmt.Println(resultadoAluno(2.5))
	fmt.Println(resultadoAluno(4))
	fmt.Println(resultadoAluno(7))
	interruptorLampada()
	interruptorLampada()

}
