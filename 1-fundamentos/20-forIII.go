package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Gerador de números aleatórios
	target := rand.Intn(10) + 1 // 1 e 10
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Tente adivinhar um número entre 1 e 10")

	var guess int

	for {
		fmt.Print("Digite seu palpite: (ou 0 para sair)")
		fmt.Scan(&guess)

		if guess == 0 {
			fmt.Println("Você desistiu. O número era: ", target)
			break
		}

		if guess > target {
			fmt.Println("Seu palpite é maior. Tente novamente!")
		} else if guess < target {
			fmt.Println("Seu palpite é menor. Tente novamente!")
		} else {
			fmt.Println("Parabéns, Você acertou!")
		}
	}
}
