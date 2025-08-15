package main

import "fmt"

func main() {
	notas := []float64{8.5, 9.0, 7.5, 6.8, 9.3}

	fmt.Println("Calculo da média com Slice")
	// Calculando a média
	var total float64
	for _, nota := range notas {
		total += nota // total = total + nota
	}

	media := total / float64(len(notas))
	fmt.Println("A soma das notas é: ", total)
	fmt.Println("A média das notas é: ", media)
}
