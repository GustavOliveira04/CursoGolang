package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	// Criando 3 Goroutines que vão ler do canal
	for i := 1; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for msg := range ch {
				fmt.Printf("Goroutine %d recebeu %s \n", id, msg)
				time.Sleep(time.Second * 500)
			}
		}(i)
	}

	// Enviando msg para o canal
	messages := []string{"Olá", "Mundo", "Go", "Concorrência", "é", "incrivel"}
	for _, msg := range messages {
		ch <- msg
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)

	wg.Wait()
	fmt.Println("Todas as Goroutines terminaram")
}
