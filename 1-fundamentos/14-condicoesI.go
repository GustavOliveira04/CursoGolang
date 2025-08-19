package main

import "fmt"

func main() {
	var idade int

	fmt.Println("Informe a sua idade: ")
	fmt.Scan(&idade)

	// fazendo primeira condicional (usa if, else e else if)
	if idade >= 18 {
		fmt.Println("Você é maior de idade!")
	} else {
		fmt.Println("Você é menor de idade!")
	}
}
