package main

import "fmt"

func main() {
	fmt.Println("Teste Golang!")
	fmt.Println("Utilizando IF/ELSE")
	fmt.Println(parseInteiro(2.0))
	fmt.Println(parseInteiro(1.5))
	fmt.Println(parseInteiro(1.4))
	fmt.Println(parseInteiro(0.9))
	fmt.Println(parseInteiro(0.4))
	fmt.Println(parseInteiro(0.0))
	fmt.Println(parseInteiro(-0.4))
	fmt.Println(parseInteiro(-0.5))
	fmt.Println(parseInteiro(-1.4))
	fmt.Println(parseInteiro(-1.5))
	fmt.Println(parseInteiro(-2.0))
	fmt.Println("Utilizando SWITCH/CASE")
	fmt.Println(parseInteiro2(2.0))
	fmt.Println(parseInteiro2(1.5))
	fmt.Println(parseInteiro2(1.4))
	fmt.Println(parseInteiro2(0.9))
	fmt.Println(parseInteiro2(0.4))
	fmt.Println(parseInteiro2(0.0))
	fmt.Println(parseInteiro2(-0.4))
	fmt.Println(parseInteiro2(-0.5))
	fmt.Println(parseInteiro2(-1.4))
	fmt.Println(parseInteiro2(-1.5))
	fmt.Println(parseInteiro2(-2.0))

}

// Pasre de números Reais Positivos e Negativos para inteiro com arredondamento
func parseInteiro(x float64) int { //Converte arredondando para inteiro
	a := x - float64(int(x))
	if a >= 0 {
		if a >= 0.5 { //trata números positivos
			return int(x) + 1
		} else {
			return int(x)
		}
	} else {
		if a <= -0.5 { //trata números negagtivos
			return int(x) - 1
		} else {
			return int(x)
		}
	}
}

// Mesma funcionalidade que parseInteiro porém com Switch/Case
func parseInteiro2(x float64) int { //Converte arredondando para inteiro
	a := x - float64(int(x))

	switch {
	case a > 0.0 && a < 0.5:
		return int(x)

	case a >= 0.5:
		return int(x) + 1

	case a < 0.0 && a > -0.5:
		return int(x)

	case a <= -0.5:
		return int(x) - 1

	default: //caso a = 0
		return int(x)
	}
}
