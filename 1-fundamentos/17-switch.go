package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Digite a nota do aluno (0 a 10)")
	fmt.Scan(&nota)

	// utilizando outro condicional switch (funciona quase igual no JS)
	// Cada switch possui seus casos (condições) seguidos do print e um default (else) seguido do print
	switch {
	case nota >= 9:
		fmt.Println("Classificação: A - Excelente")
	case nota >= 7:
		fmt.Println("Classificação: B - Bom")
	case nota >= 6:
		fmt.Println("Classificação: C - Regular")
	case nota >= 3:
		fmt.Println("Classificação: D - Insuficiente")
	case nota >= 0:
		fmt.Println("Classificação: F - Reprovado")
	default:
		fmt.Println("A nota esta inválida! Digite uma nota entre 0 e 10")
	}
}
