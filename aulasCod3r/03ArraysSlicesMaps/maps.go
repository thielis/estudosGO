package main

import (
	"fmt"
	"reflect"
)

func iterandoCadastroCliente(cadastros map[int]string) {
	//Iterando nas chaves/valor
	for chave, valor := range cadastros {

		fmt.Printf("\nNome: %s | CPF: %d\n", valor, chave)

	}
}

func iterandoEmplacamento(cadastros map[string]string) {
	//Iterando nas chaves/valor
	for chave, valor := range cadastros {

		fmt.Printf("\nCarro: %s | Placa: %s\n", valor, chave)

	}
}

func iterandoCadastroCarro(cadastros map[string]map[string]int) {
	//Iterando nas chaves/valor
	for placa, carro := range cadastros {

		fmt.Println("Placa:", placa, carro)

	}
}

func main() {
	//MAPS são estruturas Chave x Valor muito úteis para definir modelos etc
	//Forma de iniciar um MAP zerado (sem valor ou nulo - nil)

	CadastroCliente := make(map[int]string)
	fmt.Println("Tipo de Variável:", reflect.TypeOf(CadastroCliente))

	//Inserindo cadastros
	CadastroCliente[51758660015] = "Maria"
	CadastroCliente[28341849097] = "Julio"
	CadastroCliente[92517565031] = "Joao"
	CadastroCliente[68693806009] = "Pedro"

	//Imprimindo MAP completo
	fmt.Println(CadastroCliente)
	// Consultando nome pelo CPF
	fmt.Println(CadastroCliente[28341849097])
	// Excluindo cadastro
	delete(CadastroCliente, 68693806009)
	fmt.Println(CadastroCliente)
	CadastroCliente[68693806009] = "Pedro"

	//Iterando nas chaves/valor
	iterandoCadastroCliente(CadastroCliente)

	//Suprimindo CPF
	for _, nome := range CadastroCliente {

		fmt.Printf("\nNome: %s\n", nome)

	}
	//Suprimindo Nome
	for cpf, _ := range CadastroCliente {

		fmt.Printf("\nCPF: %d\n", cpf)

	}
	fmt.Printf("\n-----------------\n")

	//se remover o "_" e colocar a variável direto sempre pegará a primeira chave
	for chave := range CadastroCliente {

		fmt.Println(chave)

	}

	//Declarando o MAP direto sem inicializar
	emplacamentoCarro := map[string]string{
		"ODD-4548": "Siena",
		"HGF-8B62": "Palio",
		"MPK-8213": "Kadett",
	}

	iterandoEmplacamento(emplacamentoCarro)

	//MAP aninhado
	cadastroCarro := map[string]map[string]int{
		"ODD-4548": {
			"Siena": 2012,
		},
		"HGF-8B62": {
			"Palio": 2007,
		},
		"MPK-8213": {
			"Kadett": 1997,
		},
	}
	iterandoCadastroCarro(cadastroCarro)
}
