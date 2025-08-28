package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

// Passando infos do estoque por valor
func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

// Método pára add estoque
func (e *Estoque) AddItem(item models.Item) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID: %d] possui uma quantidade menor ou insuficiente (0 ou negativo)", item.ID)
	}
	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}
	e.items[strconv.Itoa(item.ID)] = item

	e.logs = append(e.logs, models.Log{
		TimeStamp: time.Now(),
		Action:    "Entrada de Item no Estoque",
		User:      "Admin",
		ItemId:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos items ao estoque",
	})

	return nil
}

// Método para listar items do Estoque
func (e *Estoque) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}
	return itemList
}

func (e *Estoque) ViewLogs() []models.Log {
	return e.logs
}

func (e *Estoque) CalculateTotalCost() float64 {
	var totalCost float64
	for _, item := range e.items {
		totalCost += float64(item.Quantity) * item.Price
	}
	return totalCost
}

func FindBy[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item foi encontrado")
	}
	return result, nil
}
