package main

import "fmt"

func main() {
	// Criando um slice vazio de inteiros
	var numeros []int

	// Adicionando elemtos no slice
	numeros = append(numeros, 10)
	numeros = append(numeros, 20, 30, 40)
	numeros = append(numeros, 50)

	fmt.Println("Slice de Números: ", numeros)
}
