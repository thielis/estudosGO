package main

import (
	"fmt"
	"time"
)

// Switch Case
func resultadoAlunoSwitch(nota float64) string {
	var n = int(nota)
	switch n {
	case 10:
		fallthrough //imprime A pois pula a 10 entrando direto na instrução da 9
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "valor invalido"
	}
}

// Outra forma de utilizar o switch
func resultadoAlunoSwitchBoleano(nota float64) string {
	switch {
	case nota <= 10 && nota >= 9:
		return "A"
	case nota < 9 && nota >= 7:
		return "B"
	case nota < 7 && nota >= 5:
		return "C"
	case nota < 5 && nota >= 3:
		return "D"
	case nota < 3 && nota >= 0:
		return "E"
	default:
		return "valor invalido"
	}
}

func saudacaoHoraria() {
	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("Bom dia!")
	case t < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int, int8, int16, int32, int64, uint, byte, uint16, uint32, uint64:
		return "inteiro"
	case string:
		return "string"
	case float32, float64:
		return "float"
	case func():
		return "função"
	default:
		return "não compatível com a função!"

	}
}

func main() {
	fmt.Println("Utilizando Switch Case")
	fmt.Println(resultadoAlunoSwitch(-0.5))
	fmt.Println(resultadoAlunoSwitch(-2.5))
	fmt.Println(resultadoAlunoSwitch(10))
	fmt.Println(resultadoAlunoSwitch(9))
	fmt.Println(resultadoAlunoSwitch(7))
	fmt.Println(resultadoAlunoSwitch(6))
	fmt.Println(resultadoAlunoSwitch(3))
	fmt.Println(resultadoAlunoSwitch(1))

	fmt.Println("Utilizando Switch Case Boleano")
	fmt.Println(resultadoAlunoSwitchBoleano(-0.5))
	fmt.Println(resultadoAlunoSwitchBoleano(-2.5))
	fmt.Println(resultadoAlunoSwitchBoleano(10))
	fmt.Println(resultadoAlunoSwitchBoleano(9))
	fmt.Println(resultadoAlunoSwitchBoleano(7))
	fmt.Println(resultadoAlunoSwitchBoleano(6))
	fmt.Println(resultadoAlunoSwitchBoleano(3))
	fmt.Println(resultadoAlunoSwitchBoleano(1))
	saudacaoHoraria()

	fmt.Println("Utilizando Switch Case tipo Interface")
	fmt.Println(tipo(1))
	fmt.Println(tipo(1.0))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo("teste"))
	fmt.Println(tipo(time.Now()))

}
