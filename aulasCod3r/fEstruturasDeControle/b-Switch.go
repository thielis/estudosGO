package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("\nUtilizando IF/ELSE:\n")
	fmt.Println(resultadoLetra1(9.0))
	fmt.Println(resultadoLetra1(7.0))
	fmt.Println(resultadoLetra1(6.0))
	fmt.Println(resultadoLetra1(4.0))
	fmt.Println(resultadoLetra1(2.0))
	fmt.Println(resultadoLetra1(-1.0))
	fmt.Printf("\nUtilizando Switch/Case:\n")
	fmt.Println(resultadoLetra2(9.0))
	fmt.Println(resultadoLetra2(7.0))
	fmt.Println(resultadoLetra2(6.0))
	fmt.Println(resultadoLetra2(4.0))
	fmt.Println(resultadoLetra2(2.0))
	fmt.Println(resultadoLetra2(-1.0))
	fmt.Printf("\nUtilizando Switch com incialização:\n")
	resultadoLetraAleatorio()
	resultadoLetraAleatorio()
	resultadoLetraAleatorio()
	resultadoLetraAleatorio()
	resultadoLetraAleatorio()
	fmt.Println("Switch analisando Tipos:")
	tipoVariavel(func() {})
	tipoVariavel("texto")
	tipoVariavel(1)
	tipoVariavel(1.2)
}

func resultadoLetra1(nota float64) string {
	if nota >= 8.5 {
		return "A"
	} else if nota < 8.5 && nota >= 6.5 {
		return "B"
	} else if nota < 6.5 && nota >= 5.5 {
		return "C"
	} else if nota < 5.5 && nota >= 4.0 {
		return "D"
	} else if nota < 4.0 && nota >= 0.0 {
		return "E"
	} else {
		return "Nota Inválida"
	}
}

func resultadoLetra2(nota float64) string {
	switch {
	case nota >= 8.5:
		return "A"
	case nota < 8.5 && nota >= 6.5:
		return "B"
	case nota < 6.5 && nota >= 5.5:
		return "C"
	case nota < 5.5 && nota >= 3.0:
		return "D"
	case nota < 3.0 && nota >= 0.0:
		return "E"
	default:
		return "Nota Inválida"
	}
}

// Switch com Inicialização
func resultadoLetraAleatorio() {
	notaAleatoria := rand.Float64() * 10

	switch nota := notaAleatoria; { //Bloco de inicialização quando desejado evitando números negativos
	case nota >= 9 && nota <= 10:
		fmt.Printf("Resultado A, nota: %.1f\n", nota)
	case nota < 9 && nota >= 7.5:
		fmt.Printf("Resultado B, nota: %.1f\n", nota)
	case nota < 7.5 && nota >= 6.5:
		fmt.Printf("Resultado C, nota: %.1f\n", nota)
	case nota < 6.5 && nota >= 5:
		fmt.Printf("Resultado D, nota: %.1f\n", nota)
	case nota < 5 && nota >= 0:
		fmt.Printf("Resultado E, nota: %.1f\n", nota)
	default:
		fmt.Printf("Nota inválida") //difilmente ocorrerá pois o bloco de inicialização está filtrando.
	}
}

//Switch com interfaces

func tipoVariavel(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Tipo: Inteiro")
	case float64, float32:
		fmt.Println("Tipo: Real")
	case string:
		fmt.Println("Tipo: String")
	case func():
		fmt.Println("Tipo: Função")
	default:
		fmt.Println("Tipo: Desconhecido")
	}
}
