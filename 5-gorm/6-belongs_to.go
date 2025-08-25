package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Configuração do ambiente
type Category struct {
	ID   int `gorm: "primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm: primaryKey`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Inserindo categoria
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// // Inserindo Produto
	db.Create(&Product{
		Name:       "Mesa",
		Price:      1500.00,
		CategoryID: category.ID,
	})

	// Listando produtos
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, "|", product.Category.Name, "|", product.Price)
	}
}
