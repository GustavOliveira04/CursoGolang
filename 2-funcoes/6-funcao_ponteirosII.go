package main

import "fmt"

// Criando struct
type Cliente struct {
	Nome  string
	Idade int
}

// função-1: Recebe um ponteiro para alterar os dados do cliente
func atualizarCliente(c *Cliente, novoNome string, novaIdade int) {
	c.Nome = novoNome
	c.Idade = novaIdade
}

// função principal
func main() {
	cliente := Cliente{Nome: "Gustavo", Idade: 21}
	fmt.Println("Antes da alteração: ", cliente)

	atualizarCliente(&cliente, "Gabriel", 22)
	fmt.Println("Após a alteração: ", cliente)
}
