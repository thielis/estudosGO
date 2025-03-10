package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Slices (uma parte do array, ou referecia um pedaço de um array)
	//Tem a vantagem de não precisar de tamanho definido
	//Aumenta seu tamanho de acordo com a demanda
	//Muito útil em casos que o tamanho de armazenamento de array não está bem definido
	//Não é necessário existir um array para se referenciar pode ser criado do zero
	fmt.Printf("\n----------------------\n")
	fmt.Printf("\nSlices\n")
	var pedaco []int //declarando slice
	fmt.Println(pedaco)
	pedaco = []int{1, 2, 3, 4}
	fmt.Println(pedaco)
	fmt.Println(pedaco[0])
	fmt.Println(pedaco[len(pedaco)-1])

	slice1 := []string{"a", "b", "c", "d"} //declaração curta do slice (do zero) sem referenciar outro array.
	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(slice1[len(slice1)-1])

	valores := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Tamanho da variável 'valores':", len(valores))
	slice2 := valores[0:4] // declarando um slice referenciando um array no caso "valores" obtendo do índice 0 até 4
	slice3 := valores[:4]  // declarando até o índice 4
	slice4 := valores[4:9] // declarando do índice 5 até 8
	slice5 := valores[4:]  // declarando do índice 5 em diante
	slice6 := valores[2:6] // declarando do índice 2 até 6
	slice7 := slice5[:3]   // declarando um silice do slice 5
	fmt.Println(slice2)    // valor 5 não saiu pois está na posição 4 para sair é necessário parar na posição 5
	fmt.Println(slice3)
	fmt.Println(slice4) // valor 9 saiu pois está na posição 8 e o slice percorre até a posição 9
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println("Apresentando tipo de variável 'slice2':", reflect.TypeOf(slice2))
	fmt.Println(valores[8]) // valor 9
	fmt.Println("Slice do Slice:", slice7)

	//Criando slice com o método/função make

	slicemake := make([]int, 5, 10)                                                  //neste caso será criado um slice tamanho 5 com eles zerados mas com capacidade 10
	fmt.Println("Dados so slice criado:", slicemake, len(slicemake), cap(slicemake)) //cap() mostra a capacidade do array que o slice aponta/referencia
	slicemake = append(slicemake, 1, 2, 3, 4, 5, 6)                                  //utlizado o metod append() para inserir o próprio slice mais alguns valores
	fmt.Println("Dados so slice modificado:", slicemake, len(slicemake), cap(slicemake))
	// ao estourar capacidade do slice, ele recria o um array com o dobro da capacidade, é uma padrão de GO
	// a sintaxe do appen(slice inicial, elementos ou outro slice)

	//Para mostrar que um slice tem um paradigma parecido com de um ponteiro
	//atribui-se o endereço de memória de um elemento do array valores e o esmo para o slice2 que o referencia
	//Assim será observado nos prints a seguir que o endereço é o mesmo para ambos.
	p1 := &valores[0]
	p2 := &slice2[0]

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("Antes de alterar o indice 0:", valores)
	slice2[0] = 2 //ou poderia alterar direto no endereço de memória *p2 = 2
	fmt.Println("Depois de alterar o indice 0:", valores)

}
