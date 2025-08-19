package main

import "fmt"

func main() {
	// declarando um array de 3 strings com valores iniciais
	var frutas = [3]string{"Maçã", "Banana", "Laranja"}

	fmt.Println("Array de frutas: ", frutas[0:2])
}
