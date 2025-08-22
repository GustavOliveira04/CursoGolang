package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Configuração do ambiente
type Product struct {
	ID    int `gorm: primaryKey`
	Name  string
	Price float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Select de produto (um de cada vez)
	var product Product
	db.First(&product, 1) // select do produto com o id 1

	// Select de produto pelo nome (um de cada vez)
	db.First(&product, "name = ?", "Fogão")

	// Select de produto pelo preço (um de cada vez)
	db.First(&product, "price = ?", "3500")

	// Select de todos os produtos
	var products []Product
	db.Find(&products)
	db.Limit(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println(product)

	// Select com WHERE
	var products []Product
	db.Where("price > ?", 2000).Find(&products)      // selesct pelo preço
	db.Where("name LIKE ?", "%book").Find(&products) // select pelo nome
	for _, product := range products {
		fmt.Println(product)
	}

}
