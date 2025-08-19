package main

import "fmt"

func main() {
	frutas := []string{"Maçã", "Banana", "Laranja", "Uva", "Morango"}

	// Criando um subslice, pegando do indice de 1 a 3
	// ele inicia na posição 1 (banana) e termina da posição 4 (uva)
	subslice := frutas[1:4]

	fmt.Println("Subslice de frutas: ", subslice)

	// Trocando a fruta Banana por Manga
	subslice[0] = "Manga"

	fmt.Println("Subslice: ", subslice)
	fmt.Println("Subslice de frutas atualizado: ", subslice)
}
