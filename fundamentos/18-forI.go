package main

import "fmt"

func main() {
	nomes := [4]string{"Anna", "Carlos", "Gabrielly", "Gustavo"}
	fmt.Println("Estrutura de Repetição For")

	for id, nome := range nomes {
		// fmt.Println("Id: ", id)
		// fmt.Println("Nome: ", nome)
		fmt.Println(id+1, "->", nome)
		if len(nome) > 5 {
			fmt.Println(nome, "Tem mais de 5 letras")
		}
	}
}
