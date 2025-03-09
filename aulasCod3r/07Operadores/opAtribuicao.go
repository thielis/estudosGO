package main

import "fmt"

func main() {
	//Os operadores de atribuição são muito utilizados em loops

	i := 2 //inferência de Tipo (também chamado declaração curta de variável)
	fmt.Println(i)
	i += 2 //equivale a i = i + X
	fmt.Println(i)
	i -= 2 //equivale a i = i - X
	fmt.Println(i)
	i /= 2 //equivale a i = i / X
	fmt.Println(i)
	i *= 2 //equivale a i = i * X
	fmt.Println(i)
	i %= 2 //equivale a i = i % X (resto da divisão)
	fmt.Println(i)
}
