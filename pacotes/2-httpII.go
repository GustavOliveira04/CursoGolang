package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	URL := "https://jsonplaceholder.typicode.com/todos/2"
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição: ", err)
		return
	}

	defer resp.Body.Close()

	//Usando io.ReadAll para ler o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta: ", err)
		return
	}

	fmt.Println("Resposta da API: ", string(body))
}
