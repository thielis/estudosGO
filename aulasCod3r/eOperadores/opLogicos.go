package main

import "fmt"

// Os Operadores Lógicos só funcionam com variáveis do tipo Boleana
func main() {
	var a, b, resultado bool
	//Operador lógico AND ou E
	a, b = true, true
	resultado = a && b
	fmt.Printf("Operador lógico 'E' ou 'AND'para a=%t e b=%t temos: %t\n", a, b, resultado)

	b = false
	resultado = a && b
	fmt.Printf("Operador lógico 'E' ou 'AND'para a=%t e b=%t temos: %t\n", a, b, resultado)

	//Operador OU
	a, b = true, false
	resultado = a || b
	fmt.Printf("Operador lógico 'OU' ou 'OR'para a=%t e b=%t temos: %t\n", a, b, resultado)
	a = false
	resultado = a || b
	fmt.Printf("Operador lógico 'OU' ou 'OR'para a=%t e b=%t temos: %t\n", a, b, resultado)

	//Operador "OU Exclusivo" ou "Diferente de"
	a, b = true, false
	resultado = a != b
	fmt.Printf("Operador lógico 'OU' ou 'OR'para a=%t e b=%t temos: %t\n", a, b, resultado)
	a = false
	resultado = a != b
	fmt.Printf("Operador lógico 'OU' ou 'OR'para a=%t e b=%t temos: %t\n", a, b, resultado)

	//Operanor "NÃO" ou "NOT"
	a, b = true, false
	resultado = !a
	fmt.Printf("Operador lógico 'NAO' ou 'NOT'para a=%t temos: %t\n", a, resultado)
	resultado = !b
	fmt.Printf("Operador lógico 'NAO' ou 'NOT'para b=%t temos: %t\n", b, resultado)

	celular, memoriaSD, corote := Compras(false, false)
	fmt.Printf("Compras:\nCelular: %t\nMemoria SD: %t\nCorote: %t\nCuidar do fígado: %t\n", celular, memoriaSD, corote, !corote)

}

//Exemplo de uma função com operadores lógicos

func Compras(servico1, servico2 bool) (bool, bool, bool) {
	cel := servico1 && servico2    //AND
	memSD := servico1 != servico2  //OU Exclusivo (Operador Unário) / Diferente
	chapar := servico1 || servico2 //OR
	return cel, memSD, chapar
}
