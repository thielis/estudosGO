package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	//conversões de tipos
	num1 := 0.2
	num1int := int(num1)
	fmt.Printf("Valor de num1: %.2f, convertido para int: %d\n", num1, num1int)
	num10 := 1.7
	num10int := int(num10)
	fmt.Printf("Valor de num1: %.2f, convertido para int: %d\n", num10, num10int)
	numstring := fmt.Sprint(num1) // converte em string
	num2 := 97
	num3 := 3.3
	string1 := "Palavr"
	string2 := "Valor é: "
	num2float := float64(num2)
	boleanoString := "false"

	//Como GO é fortemente tipado é necessário fazer
	//conversões para operar sempre com elementos
	//de mesmo tipo
	fmt.Println(string2 + numstring)
	fmt.Println(num1int + num2)
	fmt.Println(num2float + num3)

	//A função string([valor int]) converterá o valor inteiro
	//para o caracter representado pelo código da tabela ASCII
	fmt.Println(string1 + string(num2))

	//Para converter em string é necessário o fmt.Sprint([valor])
	// Ex.:
	fmt.Println(string2 + fmt.Sprint(num2))

	//Para converter um inteiro em string utilize o método a seguir
	fmt.Println("Inteiro para String: " + strconv.Itoa(num1int))
	numstringint := "1234"
	//Converter uma string de números em inteiro (int)
	num2int, _ := strconv.Atoi(numstringint) //utiliza-se _ para ocultar a saída de err
	fmt.Printf("String para inteiro: %d, confirmando tipo é '%s'\n", num2int, reflect.TypeOf(num2int))

	//Uma conversão de uma string para boleano través de um Parse
	boleano1, _ := strconv.ParseBool(boleanoString)
	if boleano1 {
		fmt.Printf("O Boleano é Verdadeiro, ou seja '%t'\n", boleano1)
	} else {
		fmt.Printf("O Boleano é Falso, ou seja '%t'\n", boleano1)
	}

}
