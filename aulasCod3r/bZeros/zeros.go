package main

import "fmt"

func main() {

	var a int
	var b float64
	var c bool
	var d string
	var e *int
	var f rune

	fmt.Printf("#Zeros para variáveis\nInteiro: %v\nFloat64: %v\nBoleano: %v\nString: %v\nPonteiro: %v\nRuna: %v\n", a, b, c, d, e, f)
	//nil representa nulo
	//para string fica um espaço em branco ou vazio,
	//utilizando %q mostra a string vazia de forma mais visual

	fmt.Printf("Zero de uma string é: %q", d)
}
