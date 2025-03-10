package main

import "fmt"

//Após estudar os arrays e slices é importante detalhar algumas caracteristicas
//do funcionamento do append e apresentar o copy

func main() {
	// vamos declarar um array inicialmente e um slice vazio
	array1 := [...]int{1, 2, 3}
	var slice1 []int
	slice1 = array1[:]

	slice2 := make([]int, 3)
	fmt.Println(slice1, slice2)
	// ao utilizar append em um slice que não foi iniciado
	// os valores inseridos a partir de outro slice ou valores informados
	// ocuparão de forma sequencial a partir do índice 0
	// o método append só aceita valores diretos ou slices, não aceita outro array o mesmo para copy

	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	//Obs: se o slice que receber a cópia não estiver definido, for apenas nulo não funcionará neste caso só o append
	//se o tamanho so slice que for receber a cópia for menor que e a quantidade de elementos só copiará até onde o slice está definido
	//em resumo o append aumenta o tamanho do slice de forma automática o copy não.

	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice1, slice2)

	//Exemplo de copy com slice menor
	slice3 := make([]int, 2)
	copy(slice3, slice2)
	fmt.Println(slice2, slice3)

	//Copiando definido os elementos, mas sobrescreve os que estavam,
	//ou seja, altera as memórias de referência.
	copy(slice3, slice2[3:5])
	fmt.Println(slice2, slice3)
}
