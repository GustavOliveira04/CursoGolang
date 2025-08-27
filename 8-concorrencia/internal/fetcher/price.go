package fetcher

import (
	"math/rand"
	"time"
)

// Funcão buscar preços em diferentes sites
// Simulação de um site que demora 1 segundo

func FetchPriceFromSite1() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

// Simulação de um site que demora 2 segundo
func FetchPriceFromSite2() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}

// Simulação de um site que demora 3 segundo
func FetchPriceFromSite3() float64 {
	time.Sleep(2 * time.Second)
	return rand.Float64() * 100
}
