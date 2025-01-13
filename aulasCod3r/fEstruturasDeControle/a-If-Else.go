package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nota1, nota2, nota3, notaInvalida := 8.5, 6.5, 4.0, -6.0
	nome1, nome2 := "Thielis", "Joao"

	fmt.Printf("\nNota: %.1f, Status: %s, %s\n", nota1, resultadoIfElse(nota1), resultadoCategorico(nota1))
	fmt.Printf("\nNota: %.1f, Status: %s, %s\n", nota2, resultadoIfElse(nota2), resultadoCategorico(nota2))
	fmt.Printf("\nNota: %.1f, Status: %s, %s\n", nota3, resultadoIfElse(nota3), resultadoCategorico(nota3))
	fmt.Printf("\nResultados Aleatórios:\n")
	aprovacaoAleatoria()
	aprovacaoAleatoria()
	aprovacaoAleatoria()
	aprovacaoAleatoria()
	aprovacaoAleatoria()
	validacoes(nota1)
	validacoes(nota3)
	validacoes(notaInvalida)
	informacoes(nome1)
	informacoes(nome2)
}

//Exemplos de IF, porém inclusos como uma função

// IF/ELSE
// Esta função que contém o IF/ELSE retorna String
func resultadoIfElse(nota float64) string {
	if nota >= 5.0 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

// IF/ELSE IF/ELSE
// Útil quando necessário avaliar um valor intermediário
// porém para várias casos podr ser interessante utilizar Switch/Case

func resultadoCategorico(nota float64) string {
	if nota >= 8.5 {
		return "Rendimento Alto"
	} else if nota < 8.5 && nota >= 6.5 {
		return "Rendimento Médio"
	} else {
		return "Rendimento Baixo"
	}
}

// IF com Inicialização (init)
// É muito útil para filtrar resultados dentro de uma função,
// Ou inicializar um valor de acordo com contido em uma variável/ponteiro,etc
func aprovacaoAleatoria() {
	floatAleatorio := rand.Float64() * 10
	if i := floatAleatorio; i >= 5 { // bloco de inicialização do IF
		fmt.Printf("Aprovado, nota: %.1f\n", i)
	} else {
		fmt.Printf("Reprovado, nota: %.1f\n", i)
	}

}

//É possível utilizar vários IFs seguidos porém quando
//rodar o código o fluxo passará por cada opção neste caso
//geralmente é mais interessante utilizar o Switch/Case
//Caso sejam várias validações aí pode-se utilizar ambos

// IF seguidos, útil mara multivalidações de uma variável
func validacoes(nota float64) {
	if nota >= 5 {
		fmt.Printf("\nNota: %.1f, Status: Aprovado\n", nota)
	}
	if nota <= 5 && nota >= 0 {
		fmt.Printf("\nNota: %.1f, Status: Reprovado\n", nota)
	}
	if nota < 0 {
		fmt.Printf("\nNota: %.1f, Valor Inválido(Negativo)\n", nota)
	}

}

func informacoes(nome string) {
	if len(nome) > 6 {
		fmt.Printf("\nNome %s é grande contem %d letras.\n", nome, len(nome))
	}
	if len(nome) <= 6 {
		fmt.Printf("\nNome %s é normal contem %d letras.\n", nome, len(nome))
	}
	if nome == "Thielis" {
		fmt.Println("Nome OK")
	} else {
		fmt.Println("Nome diferente")
	}

}
