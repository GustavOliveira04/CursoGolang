package main

import (
	"fmt"
	"strings"
)

func main() {

	movieName := "Top Gun"
	movieName2 := "top Gun"

	fmt.Println((movieName == movieName2))

	movieDescription := `
Top Gun Maveric é um filme de aviação e 
aventura muito consagrado 
na industria
	`
	// convertendo texto para maiusculo (use ToUpper)
	fmt.Println(strings.ToUpper(movieDescription))

	// convertendo texto para minusculo (use ToLower)
	fmt.Println(strings.ToLower(movieDescription))

	// primeira letra de cada palavra em maiusculo (use Title)
	fmt.Println(strings.Title(movieDescription))

	// encontrar a posição de um caráctere (use Index)
	fmt.Println(strings.Index(movieName, "p"))

	// contando o numero de ocorrencias de um caráctere (use Count)
	fmt.Println(strings.Count(movieDescription, "a"))
	fmt.Println(strings.Count(movieDescription, "e"))

	// substituir um elemento por outro (use RepleceAll para q ele mude na string TODA)
	fmt.Println(strings.ReplaceAll(movieDescription, "filme", "série"))

}
