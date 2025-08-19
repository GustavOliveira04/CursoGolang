package main

import "fmt"

// função-1: Nome completo
func fullName(firstName, LastName string) {
	fmt.Printf("Nome completo: %s %s \n", firstName, LastName)
}

// função-2: Soma de números
func summNumbers(a, b int) int {
	return a + b
}

// função-3: Endereço
func address(country string) {
	if country == "" {
		country = "Brasil"
	}
	fmt.Printf("Eu moro no %s\n", country)
}

// função principal
func main() {
	fmt.Println("Utilizando função com parametros (argumentos)")
	fullName("Gustavo", "Freitas")
	fmt.Printf("Soma: %d \n", summNumbers(100, 103))
	address("")
	address("Canadá")
}
