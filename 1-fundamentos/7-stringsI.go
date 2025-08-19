package main

import (
	"fmt"
	"strings"
)

func main() {
	// Go é case sensitive!!

	movieName := "Top Gun"
	movieName2 := "top Gun"

	fmt.Println((movieName == movieName2))

	// Usar o ` (crase) para escrever uma string em mais de uma linha.
	movieDescription := `
Top Gun Maveric é um filme de aviação e 
aventura muito consagrado 
na industria
	`
	// comando line para repetir algo na linha um tanto de vezes
	line := "="
	fmt.Println(strings.Repeat(line, 40))
	// essa seria a saia dessa parte: =========================================
	fmt.Println(movieDescription)

	// Verifica se uma palavra existe dentro de uma string (use Contains)
	fmt.Println(strings.Contains(movieDescription, "top"))
	fmt.Println(strings.Contains(movieDescription, "filme"))

}
