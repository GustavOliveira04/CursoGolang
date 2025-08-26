package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	// Teste para o valor de 1000
	tax := CalculateTax(1000.0)
	assert.Equal(t, 10.0, tax, "O imposto para 1000 deve ser 10")

	// Teste para o valor 0
	tax = CalculateTax(0.0)
	assert.Equal(t, 0.0, tax, "O imposto para 0 deve ser 0")

	// Teste para valores negativos
	tax = CalculateTax(-100.0)
	assert.Equal(t, 0.0, tax, "O imposto para -100 deve ser 0")

	// Teste para valor menor q 1000 e maior q 0
	tax = CalculateTax(500.0)
	assert.Equal(t, 5.0, tax, "O imposto para 500 deve ser 0")
}
