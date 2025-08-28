package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	// Titulo do script
	fmt.Println("Gerenciamento de Estoque:")
	items := []models.Item{
		{ID: 2, Name: "Camisa", Quantity: 10, Price: 50},
		{ID: 3, Name: "Fone", Quantity: 4, Price: 20},
		{ID: 4, Name: "Mouse", Quantity: 3, Price: 30},
		{ID: 5, Name: "Cabo", Quantity: 3, Price: 15},
	}
	estoque := services.NewEstoque()

	for _, item := range items {
		err := estoque.AddItem(item)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	fmt.Println(estoque.ListItems())
	for _, item := range estoque.ListItems() {
		fmt.Printf("\nID: %d | Item: %s | Quantidade: %d | Preço: %.2f",
			item.ID, item.Name, item.Quantity, item.Price)
	}

	fmt.Println()

	// Exibir o Log com mais informações
	logs := estoque.ViewLogs()
	for _, log := range logs {
		fmt.Printf("\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s",
			log.TimeStamp.Format("02/01 15:04:05"), log.Action, log.User, log.ItemId, log.Quantity, log.Reason)
	}

	fmt.Println("\nValor Total R$:", estoque.CalculateTotalCost())

	searchItem, err := services.FindBy(items, func(item models.Item) bool {
		// Filtragem por Nome
		return item.Name == "Camisa"
		// return item.Name == "Ps5"

		// Filtragem por Preço
		// return item.Price <= 30
		// return item.Price > 50

		// Filtragem por Quantidade
		// return item.Quantity >= 5
		// return item.Quantity <= 5

		// Filtragem por ID
		// return item.ID < 4
		// return item.ID > 4
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(searchItem)
}
