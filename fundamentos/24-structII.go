package main

import "fmt"

// Definindo a struct Pessoa (O struct é definido antes da func main!!!)
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco string
}

func main() {
	// Criar a instância da struct Pessoa
	pessoa1 := Pessoa{
		Nome:     "Gustavo Freitas",
		Idade:    21,
		Endereco: "Rua das Oliveiras, 650",
	}
	fmt.Printf("Informações da Pessoa: \n")
	fmt.Printf("Nome: %s \n", pessoa1.Nome)
	fmt.Printf("Idade: %d \n", pessoa1.Idade)
	fmt.Printf("Endereço: %s \n", pessoa1.Endereco)
}
