package main

import "fmt"

func main() {
	var nome string
	var idade int
	var altura float64
	var maiorDeIdade bool

	fmt.Println("Informe o seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Informe sua idade:")
	fmt.Scan(&idade)

	fmt.Println("Informe sua altura:")
	fmt.Scan(&altura)

	maiorDeIdade = idade >= 18

	// %s funciona como um placeholder. Vai exibir uma mensagem no resultado
	// %s para strings
	// %d para inteiros
	// %f para floats
	// %v para booleanos

	// para trabalhar com tipos deve se usar o Printf e não o Println!!

	fmt.Printf("\nDados Pessoais:\n")
	fmt.Printf("Nome: %s \n", nome)
	fmt.Printf("Idade: %d \n", idade)

	// .2 após o percent para limitar o float a apenas duas casas decimais
	fmt.Printf("Altura: %.2f \n", altura)
	fmt.Printf("Maior de Idade: %v  \n", maiorDeIdade)
}
