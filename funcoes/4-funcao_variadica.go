package main

import (
	"fmt"
)

// função-1: Soma de números com variádicas
func sum(nums ...int) {
	sumTotal := 0
	for _, n := range nums {
		sumTotal += n
	}
	fmt.Printf("Soma é: %d\n", sumTotal)
}

// função-2: Apresentação de cursos com variádicas
func presentation(data map[string]string) {
	for key, value := range data {
		fmt.Printf("%s - %s\n", key, value)
	}
}

// função principal
func main() {
	sum(7)
	sum(7, 9)
	sum(7, 9, 8)
	sum(7, 9, 8, 10)
	sum(7, 9, 8, 10, 6)

	presentation(map[string]string{
		"name":     "Python",
		"category": "Back-end",
		"level":    "Iniciante",
	})

	presentation(map[string]string{
		"name":     "Visão Computacional com Python",
		"category": "IA",
		"level":    "Avançado",
	})

	presentation(map[string]string{
		"name":     "Dashboards com dash",
		"category": "Data Science",
		"level":    "Intermediário",
	})
}
