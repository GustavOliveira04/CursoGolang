package main

import (
	"fmt"
	"strings"
)

func main() {
	// Criando ujm map para contar as palavras
	text := "go é uma linguagem de programação e go é muito rápida ela é parecida com C++"
	words := strings.Fields(text)
	fmt.Println(words)

	// é usado make para fazer um map (oq esta em [] é a chave e o de fora é o valor)
	wordCount := make(map[string]int)

	// Contagem da frequencia de palavras
	for _, word := range words {
		wordCount[word]++
	}

	// Imprimir as frequencias
	fmt.Println("Contagem de palavras")
	for word, count := range wordCount {
		fmt.Printf("Palavra: %s | Frequencia: %d \n", word, count)
	}

}
