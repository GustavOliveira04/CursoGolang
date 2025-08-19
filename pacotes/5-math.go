package main

import (
	"fmt"
	"math"
)

// 1- Acessar o número de PI
func acessarPi() {
	fmt.Printf("PI Arredondado (2 casas deciamais): %.2f\n", math.Pi)
}

// 2- Número de Euler
func acessarEuler() {
	fmt.Printf("Euler Arredondado (2 casas deciamais): %.2f\n", math.E)
}

// 3- Arredondando números para cima e para baixo
func arredondar() {
	num := 10.4
	fmt.Println("Arrendondando para cima: ", math.Ceil(num))
	fmt.Println("Arrendondando para baixo: ", math.Floor(num))
}

func main() {
	acessarPi()
	acessarEuler()
	arredondar()
	fmt.Println("Potencia de 5 elevado a 5: ", math.Pow(5, 5))
	fmt.Println("Raiz quadrada de 169: ", math.Sqrt(169))
	fmt.Println("Logaritmo de 10: ", math.Log(10))
}
