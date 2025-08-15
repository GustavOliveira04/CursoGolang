package main

import "fmt"

// Definindo a struct Carro (O struct é definido antes da func main!!!)
type Carro struct {
	Modelo string
	Ano    int
	Cor    string
}

func main() {
	// Criar a instância da struct Carro
	carro1 := Carro{
		Modelo: "Fusca",
		Ano:    1965,
		Cor:    "Verde",
	}
	fmt.Printf("Informações do Carro: \n")
	fmt.Printf("Modelo: %s \n", carro1.Modelo)
	fmt.Printf("Ano: %d \n", carro1.Ano)
	fmt.Printf("Cor: %s \n", carro1.Cor)
}
