package main

import (
	"fmt"
	"math"
	m "math" //a letra m na frente é um atalho para usar o pacote math e suas funções
	"reflect"
)

func main() {
	//Print básico no console
	fmt.Print("Teste GO!")
	fmt.Printf("\n")              //pulando uma linha
	fmt.Println("Novo Teste GO!") //print pulando uma linha após a saída

	const PI float64 = 3.1415 // declarando uma constante
	var raio = 3.2            // declarando um float

	//A seguir temos a declaração curta de variável ":=";
	//Utilizando o pacote math com "m." utilizando a função Pow para potenciação;
	//O "raio" é a base e o "2" é o expoente que vai eleva o valor do "raio" equivale a raio^2.
	area := PI * m.Pow(raio, 2)

	//Somente com o Printf é possível formatar as saídas de variáveis.
	//Abaixo tem mais abordagens com outros exemplos
	fmt.Printf("A área da circunferência é %.2f\n", area) // %.2f indica que será exibido um numero float com 2 casas depois do ponto

	//declarando constantes
	const (
		a = 1
		b = 2
	)
	//declarando variáveis numéricas inteiras
	var (
		c = 3
		d = 4
	)

	fmt.Printf("Temos as constantes a=%d e b=%d;\nTemos as variáveis inteiras c=%d e d=%d.\n", a, b, c, d) //imprimindo valores das variáveis

	var e, f bool = true, false //declarando variável boleana
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!" //declaração curta de múltiplas variáveis de tipos diferentes
	fmt.Println(g, h, i)

	//>>>Utilizando as funções print

	//A forma comentada abaixo não funciona pois é um tipo string concatenando com um número
	//fmt.Print("PI é o número: " + PI)

	//Para funcionar é necessário transformar em string primeiro, pois para concatenar os tipos
	//precisam ser iguais

	PIs := fmt.Sprint(PI)                // converte em string
	fmt.Println("PI é o número: " + PIs) // o "+" concatena apenas tipos iguais

	//Outras formas de executar sem precisar da conversão
	fmt.Println("Pula linha após exibir no console.") // o ln no final pula uma linha após exibir o texto
	fmt.Print("PI é o número:", PI)                   // o "," concatena de forma automática o valor a ser exibido

	//#- Utilizandofmt.Printf()

	//Utilizando Printf em que é possível formatar as saídas
	// %d - inteiro / %f - float / %.2f - Float com 2 casas decimais
	// %t - boleano / %s - string /
	fmt.Printf("\nInteiro: %d\nFloat: %f\nFloat com 2 casas: %.2f\nBoleano: %t\nString: %s\n", c, PI, PI, e, i)

	// %v - verbose é o modo que adapta de acordo com o tipo
	fmt.Printf("Verbose: %v, %v, %v, %v\n", c, PI, e, i)

	//Outras formas de usar o "\n" para pular linhas:
	fmt.Printf("Inteiro: %d\nFloat: %f, %.2f\nBoleano: %t\nString: %s\n", c, PI, PI, e, i)

	//Pacote reflect muito bom para retornar o tipo de variável útil para testes
	//inteiros
	fmt.Println("Variável do tipo:", reflect.TypeOf(c))

	//mesmo sendo uma constante só retorna o tipo
	fmt.Println("Variável do tipo:", reflect.TypeOf(a))

	//String
	fmt.Println("Variável do tipo:", reflect.TypeOf(i))

	//Float
	fmt.Println("Variável do tipo:", reflect.TypeOf(PI))

	//Pode ser colocado o valor direto na função reflect.TypeOf
	fmt.Println("Variável do tipo literal:", reflect.TypeOf(10))
	fmt.Println("Variável do tipo literal:", reflect.TypeOf("Texto"))
	fmt.Println("Variável do tipo literal:", reflect.TypeOf(9.25))
	fmt.Println("Variável do tipo literal:", reflect.TypeOf(true))

	//Outros tipos de variáveis numéricas
	var xa uint8 = 10
	var xb byte = 11 //O tipo "uint8" também pode ser representado pelo alias "byte"
	fmt.Printf("Este número %d e este %d, ambos são do tipo %v!\nEsse tipo de variável só tem números inteiros, postivos!\nEntão, por ser de 8 bits só é possível retornar seu máximo que é: %d\n", xa, xb, reflect.TypeOf(xb), math.MaxUint8)

	xc, xd := math.MinInt8, math.MaxInt8
	fmt.Printf("Para variável do tipo %s, inteiro de 8 bits o mínimo é %d e o máximo %d\n", reflect.TypeOf(xc), xc, xd)

	xe, xf := m.MinInt16, m.MaxInt16
	fmt.Printf("Para variável do tipo %s, inteiro de 16 bits o mínimo é %d e o máximo %d\n", reflect.TypeOf(xe), xe, xf)

	xg, xh := math.MinInt32, math.MaxInt32
	fmt.Printf("Para variável do tipo %s, inteiro de 16 bits o mínimo é %d e o máximo %d\n", reflect.TypeOf(xg), xg, xh)

	xi, xj := math.MinInt64, math.MaxInt64
	fmt.Printf("Para variável do tipo %s, inteiro de 16 bits o mínimo é %d e o máximo %d\n", reflect.TypeOf(xi), xi, xj)

	//Variável do tipo runa que é "rune" armazena o valor que representa o caractere da tabela unicode
	//É do tipo int32, e deve ser um caractere entre aspas simples: ex.: 'b'
	var ya rune = 'a'
	yastring := string(ya)
	yb := 'b' //método de declaração curta para variável rune.
	ybstring := string(yb)
	fmt.Printf("O tipo de variável rune é %v em que representa o caracter ''%s'' representado pelo número %v da tabela unicode.\n", reflect.TypeOf(ya), yastring, ya)
	fmt.Printf("O tipo de variável rune é %v em que representa o caracter ''%s'' representado pelo número %v da tabela unicode\n", reflect.TypeOf(yb), ybstring, yb)

	/*Um observação importante é o tipo Float se for feito via :=
	o tipo de variável será de acordo com a arquitetura (32 ou 64bits) do computador
	em que estiver funcionando o programa*/

	numeroFloat := 25.09
	var numeroFloat32 float32 = 25.09
	var numeroFloat64 float64 = 25.09

	fmt.Printf("Esse tipo de variável baseada no CPU é: %s ", reflect.TypeOf(numeroFloat))
	if reflect.TypeOf(numeroFloat) == reflect.TypeOf(numeroFloat32) {
		fmt.Printf("(Arquitetura 32 bits)\n")
	}
	if reflect.TypeOf(numeroFloat) == reflect.TypeOf(numeroFloat64) {
		fmt.Printf("(Arquitetura 64 bits)\n")
	}

	fmt.Printf("Esse tipo de variável é: %s\n", reflect.TypeOf(numeroFloat32))
	fmt.Printf("Esse tipo de variável é: %s\n", reflect.TypeOf(numeroFloat64))
	fmt.Printf("Utilizando negação na variável boleana temos: %t\n", !h)

	//Retornando tamanho de uma string
	fmt.Printf("Tamanho da string '%s' é: %v\n", i, len(i))

	//Outra forma de declarar a string utilizanco "`" crase
	s1 := ` Teste de 
	string 
	
	escrita 
	
	em várias linhas `

	//retorna o tamanho da string com len() porém é contabilizado os espaços em branco
	fmt.Printf("A string '%s'\n tem tamanho: %v\n", s1, len(s1))

	var z, x, v byte

	fmt.Println(z, x, v)

	//Declaração de Arrays e Slices
	// Array é quando temos um vetor com tamanho e tipo definido
	fmt.Println(">>Arrays<<")
	var numeros1 [3]int               //declara um array zerado
	var numeros2 = [3]int{1, 2, 3}    //declara um array de tamanho 3 definido os valores de indice 0 a 2
	numeros3 := [5]int{1, 2, 3, 4, 5} // declaração curta de variável definindo um array
	fmt.Println(numeros1)
	fmt.Println(numeros2)
	fmt.Println(numeros2[0]) // representa o 1º indice
	fmt.Println(numeros3)
	fmt.Println(numeros3[4]) // representa o indice 5

	fmt.Println(">>Slices<<")
	//Slice é um pedaço de array em que seu tamanho não é fixo
	//Quanto mais itens você atribui a ele mais aumenta seu tamanho
	var numeros4 []int // declara um slice zerado
	fmt.Println("Slice zerado", numeros4)
	var numeros5 = []int{1, 2, 3}   //declara um slice com 1,2,3 como valores
	numeros6 := []int{1, 2}         // declaração curta
	numeros4 = []int{1, 2, 3, 4, 5} //atribuindo valor ao slice numeros4
	fmt.Println(numeros5)
	fmt.Println(numeros6)
	fmt.Println(numeros1)
	fmt.Println(numeros1)
	fmt.Println("Após atribuir valores a variavel 'numeros4':")
	fmt.Println(numeros4)

}
