package main

import "fmt"

// Em GO não existe o Operador ternario em que entra um resultado e
// existem 2 possibilidades de resposta
func main() {
	fmt.Println("Resultado final da nota é: ", resultadoFinal(5.9)) // chamando a função inserindo um valor

}

func resultadoFinal(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}
