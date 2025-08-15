package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Digite sua nota (0 a 10)")
	fmt.Scan(&nota)

	// nota < 0 ou nota > 10
	if nota < 0 || nota > 10 {
		fmt.Println("Nota inválida! Digite uma nota entre 0 e 10")
	} else if nota >= 9 {
		fmt.Println("Excelente!, Aprovado")
	} else if nota >= 8 {
		fmt.Println("Muito bom!, Aprovado")
	} else if nota >= 6 {
		fmt.Println("Satisfatório!, Aprovado")
	} else {
		fmt.Println("Reprovado!")
	}
}
