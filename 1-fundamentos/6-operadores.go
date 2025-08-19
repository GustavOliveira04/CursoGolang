package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Informe o primeiro número")
	fmt.Scan(&num1)

	fmt.Println("Informe o segundo número")
	fmt.Scan(&num2)

	// Valores lógicos

	// + para soma
	soma := num1 + num2

	// - para subtração
	sub := num1 - num2

	// * para multiplicação
	mult := num1 * num2

	// / para divisão
	div := num1 / num2

	// % para módulo (resto da divisão)
	mod := num1 % num2

	// usando printf para mostrar um placeholder com os valores passados acima
	fmt.Printf("Soma de %d e %d é %d \n", num1, num2, soma)
	fmt.Printf("Subtração de %d e %d é %d \n", num1, num2, sub)
	fmt.Printf("Multiplicação de %d e %d é %d \n", num1, num2, mult)
	fmt.Printf("Divisão de %d e %d é %d \n", num1, num2, div)
	fmt.Printf("Resto da Divisão de %d e %d é %d \n", num1, num2, mod)

	// Atribuição
	num1 += 1 // num1 = num1 + 1
	num1 -= 1 // num1 = num1 - 1
	num1 *= 1 // num1 = num1 * 1
	num1 /= 1 // num1 = num1 / 1

	// Comparação

	// maior que...
	bigger := num1 > num2

	// menor que...
	smaller := num1 < num2

	// igual á (comparação)
	equal := num1 == num2

	// diferente de...
	different := num1 != num2

	// maior ou igual que...
	biggerEqual := num1 >= num2

	// menor ou igaul que...
	smallerEqual := num1 <= num2

	fmt.Printf("O número %d é maior que %d? %v \n", num1, num2, bigger)
	fmt.Printf("O número %d é menor que %d? %v \n", num1, num2, smaller)
	fmt.Printf("Os números %d e %d são iguais? %v \n", num1, num2, equal)
	fmt.Printf("O número %d é diferente de %d? %v \n", num1, num2, different)
	fmt.Printf("O número %d é maior ou igual a %d? %v \n", num1, num2, biggerEqual)
	fmt.Printf("O número %d é menor ou igual a %d? %v \n", num1, num2, smallerEqual)

	fmt.Println("Valor final de num1 após alterações de atribuição: ", num1)
}
