package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1 //declaração curta de variável de um inteiro

	//Ponteiro
	var p *int = nil
	/*
		Observações importantes:
		- Variável "p" representa um ponteiro do tipo int;
		- O ponteiro representa/aponta um endereço de memória;
		- Foi declarado inicialmente como "nil" (nulo) pois "p"
		  ainda não aponta para nenhum endereço, uma boa prática
		  na hora de declarar iniciando-a para garantir que não
		  está sendo referenciado a algum endereço de memória;
		- Não é possível executar operações aritiméticas com um ponteiro
		  sem antes desreferenciar o endereço para trabalhar com o valor
		  contido no endereço de memória ao qual o ponteiro representa.
	*/

	//Atribuindo um endereço de memória a um ponteiro
	p = &i
	fmt.Println("Valor de do endereço de 'i': ", &i)
	/*
		O "&" junto a variável retorna o endereço de memória da variável
		em que sempre que vamos fazer alguma operação em uma variável em
		GO é criado uma cópia do valor trabalhado em cima deste, agora
		quando trabalhamos com ponteiro estamos lidando com a referência
		ou endereço de memória
		No caso acima foi atribuído em "p" o endereço de memória que a variável "i"
		representa.
	*/
	fmt.Printf("Valor de 'i': %d\nValor que 'p' representa: %v\n", i, p)

	/*Para melhor esclarecer a idéia de referenciação observe que podemos ter o valor
	contido em p e também o endereço que ele representa: */
	fmt.Println("Valor contido em 'p': ", *p)    //desreferenciando o endereço retorna o valor com "*p"
	fmt.Println("Endereço do ponteiro 'p': ", p) //retorna o endereço de memória de p
	// Então ao executar o código observa nas saída que o endereço de "i" (&i) é igual ao de "p"

	/*
		Para executar operações com o valor contido no endereço do ponteiro "p"
		é necessário desreferenciar ou obter o valor presente neste endereçamento
		utilizando o "*" conforme a seguir:
	*/
	novoValor := *p + 1

	fmt.Println("Novo valor: ", novoValor)

	/*
		Observando as operações a seguir, observe que podemos iterar um valor na variável "i" e
		também operar sobre o valor presente na memória que "i" representa através da referência
		que o ponteiro "p" aponta.
	*/

	fmt.Println("Valor de 'i': ", i)
	i++
	fmt.Println("Valor de 'i++': ", i)
	*p++
	fmt.Println("Valor de 'i' depois de iterar sobre 'p': ", i)

	/*
		Os ponteiros são muito úteis para armazenar valores que precisam ser atualizados dentro de uma função
		e serem utilizados em outra diferente ou até processos diferentes permitindo uma flexibilidade muito
		grande dentro do código.
	*/

	*p = teste(i)
	fmt.Println("Valor de 'i' após a função 'teste' operar e salvar em 'p': ", i)
	/* O interessante do ponteiro é que ele garante que está sendo enviado e armazenado do endereço de memória
	   desejado.
	*/
	/*
		j := 10.2
		p = &j
		Observe que o j por inferência foi declarado e atribuído como um float, porém no iníco do código
		observa-se que quando a variável "p" do tipo ponteiro foi declarada, foi definido que ela só pode
		aceitar endereços de memória do tipo "int" com isso um float armazena um tamanho e tipo de dados diferentes,
		isso permite garantir que este ponteiro sempre representará um valor do tipo "int" evitando falhas de operações
		durante o código garantindo uma robustes na aplicação criada.
	*/
	//Declaração curta de um ponteiro
	g := &i
	fmt.Println(reflect.TypeOf(g), g) //retornando o tipo e endereço

	//Outro exemplo de declaração curta de um ponteiro
	f := 1.2
	h := &f
	fmt.Println(reflect.TypeOf(h), h, *h) //retornando o tipo, endereço e valor contido
}

func teste(x int) int {
	return x + 1
}
