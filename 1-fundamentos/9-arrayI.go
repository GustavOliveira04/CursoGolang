package main

import "fmt"

func main() {
	// declarando um array com 5 elementos do tipo INT
	var numeros [5]int

	// atribuindo valores as posições
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println("Array de números: ", numeros)
}
