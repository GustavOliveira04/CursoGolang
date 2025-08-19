package main

import (
	"fmt"
	"os"
)

func main() {
	// Criando um arquivo
	arquivo, err := os.Create("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo ", err)
		return
	}

	defer arquivo.Close()

	// Escrevendo no arquivo
	_, err = arquivo.WriteString("Texto de exemplo no arquivo")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo ", err)
		return
	}

	fmt.Println("Arquivo criado e texto escrito com sucesso")

	// Lendo o arquivo
	conteudo, err := os.ReadFile("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo ", err)
		return
	}

	fmt.Println("Conteúdo do arquivo: ", string(conteudo))
}
