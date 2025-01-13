package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	notafinal := []float64{8.5, 8.0, 5.5, 4.0, 2.0} // Slice
	//laço for
	//Muito utilizado para iterar em arrays/vetor/matriz como a seguir
	for i := 0; i < len(notafinal); i++ {
		fmt.Println("Nota: ", notafinal[i])
		fmt.Println(resultadoAprovacao(notafinal[i]))
		fmt.Println(resultadoCategorico(notafinal[i]))
		fmt.Println(resultadoRendimento(notafinal[i]))
	}

	//While: como não existe a clásula While em GO utiliza-se FOR para implementar um algoritimo equivalente:
	fmt.Println("While Utilizando FOR")
	i := 0       // declarando um valor inicial a variável para implementar o While
	for i < 10 { // enquanto satisfazer a condição será executado o FOR
		fmt.Println("Executando i=", i)
		i = rand.Intn(11) // aleatórios inteiros até 20
		//time.Sleep(time.Second * 2) //roda a cada 2 segundos.
	}

	variasStrings := []string{"a", "b", "c", "d", "e"}
	variosNums := []int{1, 2, 3, 4, 5}

	fmt.Println(variasStrings)
	fmt.Println(variosNums)
	varrendoArrayString(variasStrings)
	varrendoArrayInt(variosNums)
	varrendoArrayInt2(variosNums)

	//FOR como um loop infinito (muito últil para escuta de portas em um webservice por exemplo)
	for {
		fmt.Printf("\nUtilizando Loop Infinito:\n")
		fmt.Println(resultadoAprovacao(floatAleatorio()))
		fmt.Println(resultadoCategorico(floatAleatorio()))
		fmt.Printf("%.1f\n", floatAleatorio())
		fmt.Println("Para Sempres...")
		time.Sleep(time.Second) // a cada segundo ele pula para o próximo loop.
	}

}

// IF/ELSE em uma função que retorna String
func resultadoAprovacao(nota float64) string {
	if nota >= 6.5 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

// IF/ELSE IF/ELSE - utilizado quando necessário avaliar um valor intermediário

func resultadoRendimento(nota float64) string {
	if nota >= 8.5 {
		return "Rendimento Alto"
	} else if nota < 8.5 && nota >= 6.5 {
		return "Rendimento Médio"
	} else {
		return "Rendimento Baixo"
	}
}

func resultadoCategorico(nota float64) string {
	switch {
	case nota >= 8.5:
		return "A"
	case nota < 8.5 && nota >= 6.5:
		return "B"
	case nota < 6.5 && nota >= 5.5:
		return "C"
	case nota < 5.5 && nota >= 3.0:
		return "D"
	default:
		return "E"
	}
}

func floatAleatorio() float64 {
	return rand.Float64() * 10
}

func varrendoArrayString(arrayStr []string) {
	for i, arrStr := range arrayStr {
		fmt.Printf("índice %s, elemento: %d\n", arrStr, i)
	}
}

func varrendoArrayInt(arrayInt []int) {
	for arrInt := range arrayInt {
		fmt.Printf("indice: %d\n", arrInt)
	}
}

func varrendoArrayInt2(arrayInt []int) {
	for _, arrInt := range arrayInt {
		fmt.Printf("elemento: %d\n", arrInt)
	}
	//Oculta-se o índice para printar o conteúdo do slice e não o índice.
}
