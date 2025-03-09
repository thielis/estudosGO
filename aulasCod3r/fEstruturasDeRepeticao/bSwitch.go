package main

import "fmt"

// Switch Case
func resultadoAlunoSwitch(nota float64) string {
	switch nota {
	case nota <= 3.0 && nota >= 0:
		return "reprovado"

	case nota > 3.0 && nota <= 4.5:
		return "recuperação"

	case nota > 4.5:
		return "Aprovado"

	default:
		return "valor invalido"
	}
}

func main() {
	fmt.Println(resultadoAlunoSwitch(-0.5))
	fmt.Println(resultadoAlunoSwitch(2.5))
	fmt.Println(resultadoAlunoSwitch(4))
	fmt.Println(resultadoAlunoSwitch(7))

}
