package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// Iniciando uma goroutine que envia numeros para o canal
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Enviando número %d para o canal \n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	// Lendo os numeros do canal
	for num := range ch {
		fmt.Printf("Número recebido %d \n", num)
	}
}
