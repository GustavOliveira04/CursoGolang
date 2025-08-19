package main

import "fmt"

// função-1: Recebe um ponteiro como argumento e altera o valor da variavel original
func alterValue(x *int) {
	//Alterando o valor do ponteiro
	*x = *x * 2
}

// função principal
func main() {
	num := 5

	fmt.Printf("Valor inicial: %d\n", num)

	// Passando o ponteiro para a função
	alterValue(&num)
	fmt.Printf("Valor após a alteração: %d\n", num)
}
