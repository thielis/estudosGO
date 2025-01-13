package main

import (
	"fmt"
	"time"
)

func main() {
	//operadores relacionais
	a, b, c, d := 1, 1, 0, "texto"

	fmt.Printf("'%v' == '%v': %v\n", a, b, a == b)
	fmt.Printf("'%v' == '%v': %v\n", a, d, fmt.Sprint(a) == d)
	fmt.Printf("'%v' > '%v': %v\n", a, c, a > c)
	fmt.Printf("'%v' < '%v': %v\n", a, c, a < c)
	fmt.Printf("'%v' >= '%v': %v\n", a, b, a >= b)
	fmt.Printf("'%v' <= '%v': %v\n", a, b, a <= b)
	fmt.Printf("'%v' != '%v': %v\n", a, c, a != c)

	//Comparar igualdade de structs

	//Criando uma struct (não é o assunto principal deste código)
	type Teste struct {
		campo string
	}

	dados1, dados2 := Teste{campo: "teste"}, Teste{campo: "teste"}

	data1, data2 := time.Unix(0, 0), time.Unix(0, 0)

	fmt.Println("Esta igualdade é: ", data1.Equal(data2)) // para utilizar "[variável1].Equal(variável2)" é necessário ter o método definido para a struct
	fmt.Println("Esta igualdade é: ", dados1 == dados2)
}
