package main

import "fmt"

// função-1: Imprimir mensagem de boas vindas
func welcome() {
	fmt.Println("Bem Vindo ao Sistema de Filmes!")
}

// função-2: Cadastrar um filme
func createMovie() {
	var name string
	var yearRelease int
	var moviePrice float64
	var dummy string

	fmt.Println("Divite o nome do filme: ")
	fmt.Scanln(&name)

	// Limpando buffer de entrada
	fmt.Scanln(&dummy)

	fmt.Println("Divite o ano de lançamento do filme: ")
	fmt.Scanln(&yearRelease)

	fmt.Println("Divite o valor do filme: ")
	fmt.Scanln(&moviePrice)

	fmt.Printf("%s (%d) - R$ %.2f\n", name, yearRelease, moviePrice)
}

// 3- função-3: Calcular a média de notas
func calculateAverage() float64 {
	var numRatings int
	fmt.Println("Digite quantas avaliações deseja fazer para o filme: ")
	fmt.Scan(&numRatings)

	var total float64
	for i := 0; i < numRatings; i++ {
		var note float64
		fmt.Println("Digite a nota para o filme:")
		fmt.Scan(&note)
		total += note
	}

	var average float64

	if numRatings > 0 {
		average = total / float64(numRatings)
	} else {
		average = 0
	}

	return average
}

// função principal
func main() {
	welcome()
	fmt.Println("Utilizando Função")
	// createMovie()
	media := calculateAverage()
	fmt.Printf("A média das avaliações é: %.2f \n", media)
}
