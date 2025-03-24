package main

import "fmt"

// Função simples
func funcSimples() {
	//Pode ser um retorno de algo, print, processo etc
	fmt.Println("Função Simples")
}

// Função que retorna um tipo de dado sem entrada de parâmetro
func funcRetorno() string {
	//retonando uma string
	return "Teste"
}

// Função com entrada de parâmetros
// O idela é usar boleano mas isto é para nível de estudo
func funcParametros(letra1, letra2 string) int {
	//comparando letras
	if letra1 == letra2 {
		return 1
	} else {
		return 0
	}
}

func funcVariosRetornos(num1, num2 int) (string, int) {
	return "Resuntado da soma é:", num1 + num2
}

// Funções
func main() {
	funcSimples()
	funcRetorno()
	fmt.Println("Retorno é: ", funcParametros("a", "b"))
	fmt.Println("Retorno é: ", funcParametros("a", "a"))
	fmt.Println(funcVariosRetornos(1, 2))
	fmt.Println(funcVariosRetornos(2, 2))
}
