package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c, d := 2, 3, 27, 9
	//Soma
	r := a + b
	fmt.Printf("Soma de %d mais %d é igual: %d\n", a, b, r)
	//Subtração
	r = a - b
	fmt.Printf("Subtração de %d menos %d é igual: %d\n", a, b, r)
	//Multiplicação
	r = a * b
	fmt.Printf("Multiplicação de %d e %d é igual: %d\n", a, b, r)
	//Divisão
	rD := r / a
	fmt.Printf("Divisão de %d por %d é igual: %d\n", r, a, rD)
	//Resto
	r = b % b
	fmt.Printf("Resto de %d dividido por %d é igual: %d\n", b, b, r)
	r = b % a
	fmt.Printf("Resto de %d dividido por %d é igual: %d\n", a, b, r)
	//Exponenciação
	rE := math.Pow(float64(b), float64(a)) //base, expoente
	fmt.Printf("Exponeciação de base %d e expoente %d é igual: %.2f\n", b, a, rE)
	//Raiz de qualquer íncide
	rE = math.Pow(float64(c), 1.0/float64(b)) //base, Indice é igual expoente/valor
	fmt.Printf("Raiz de base %d e indice %d é igual: %.2f\n", c, b, rE)
	//Raiz quadrada
	fmt.Printf("Raiz quadrada de %d é igual: %.2f\n", d, math.Sqrt(float64(d)))

}
