package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Configuração do ambiente
type SerialNumber struct {
	ID        int `gorm: "primaryKey`
	Name      string
	ProductID int
}

type Category struct {
	ID   int `gorm: "primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm: primaryKey`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Inserindo categoria
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// // Inserindo Produto
	db.Create(&Product{
		Name:       "Mesa",
		Price:      1500.00,
		CategoryID: category.ID,
	})

	// Inserindo serial number
	db.Create(&SerialNumber{
		Number:    "123123123",
		ProductID: 1,
	})

	// Listando produtos
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, "|", product.Category.Name, "|", product.Price, "|", product.SerialNumber.Number)
	}
}
