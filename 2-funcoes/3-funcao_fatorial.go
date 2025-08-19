package main

import "fmt"

// função-1: Calcula fatorial de números
func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

// função-2: Soma os números usados no fatorial
func totalSum(num int) int {
	if num == 1 {
		return 1
	} else {
		return num + totalSum(num-1)
	}
}

// função principal
func main() {
	var number int
	fmt.Println("Digite o número para o fatorial:")
	fmt.Scan(&number)
	fmt.Printf("O fatorial de %d é %d\n", number, factorial(number))
	var num int
	fmt.Println("Digite um número para a soma:")
	fmt.Scan(&num)
	fmt.Printf("A soma total do %d é: %d\n", num, totalSum(num))
}
