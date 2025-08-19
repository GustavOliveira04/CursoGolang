package main

import "fmt"

func main() {
	// Tipo de dados inteiros (INT)
	var idade int = 30
	// var idade int = 17 caso para menor de idade

	// Tipo de dados com virgula ou ponto flutuante (FLOAT)
	var altura float64 = 1.75

	// Tipo de dado booleano (Verdaderio ou Falso)
	var maiorDeIdade bool = idade >= 18

	// Tipode dado de texto (String)
	var nome string = "Gustavo"

	fmt.Println("Dados Pessoais: ")
	fmt.Println("Nome: ")
	fmt.Println(nome)
	fmt.Println("Idade: ")
	fmt.Println(idade)
	fmt.Println("Altura: ")
	fmt.Println(altura)
	fmt.Println("Maior de Idade: ")
	fmt.Println(maiorDeIdade)

	// Para verificar o tipo de dado usando a função
	fmt.Println(fmt.Sprintf("%T", maiorDeIdade))
}
