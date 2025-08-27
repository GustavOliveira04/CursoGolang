package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	var wg sync.WaitGroup // Grupo de espera (wait group)
	wg.Add(3)             // 3 pq tem 3 Goroutines

	go func() {
		var totalPrice float64
		countPrice := 0.0
		for price := range priceChannel {
			totalPrice += price
			countPrice++
			fmt.Printf("R$ %.2f \n", price)
			fmt.Printf("Preço Médio: R$ %.2f \n", totalPrice/countPrice)
		}
	}()

	go func() {
		defer wg.Done() // Executa quando a função for executada
		priceChannel <- fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done() // Executa quando a função for executada
		priceChannel <- fetcher.FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done() // Executa quando a função for executada
		priceChannel <- fetcher.FetchPriceFromSite3()
	}()

	wg.Wait() // Aguarda todas as Goroutines terminarem
	close(priceChannel)

	fmt.Printf("\nTempo total: %s", time.Since(start))

}

// ------------------------------------------------------------------------
// Algumas anotações:

// Concorrencia ---------
// lidar com varias tarefas ao mesmo tempo. Alternância rápida entre tarefas.

// Paralelismo  ---------
// Executar várias tarefas simultareamente. Execução real em paralelo

// Goroutines   ---------
// -Threads Leves
// -Gerenciamento Automático
// -Sintaxe Simples
