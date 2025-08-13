package main

import "fmt"

func main() {
	var nome string
	var idade int
	var altura float64
	var maiorDeIdade bool

	// Scan serve como um input. Mas para usa-lo deve referenciar a variavel para atribuir um valor
	fmt.Println("Informe o seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Informe sua idade:")
	fmt.Scan(&idade)

	fmt.Println("Informe sua altura:")
	fmt.Scan(&altura)

	maiorDeIdade = idade >= 18

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(altura)
	fmt.Println(maiorDeIdade)
}
