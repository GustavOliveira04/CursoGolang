package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Configuração do ambiente
type Category struct {
	ID       int `gorm: "primaryKey"`
	Name     string
	Products []Product `gorm: "many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm: primaryKey`
	Name       string
	Price      float64
	Categories []Category `gorm: "many2many:products_categories;"`
	gorm.Model
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Inserindo categorias
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletrodomésticos"}
	db.Create(&category2)

	// // // Inserindo Produto
	db.Create(&Product{
		Name:       "Microondas",
		Price:      1000.00,
		Categories: []Category{category, category2},
	})

	// Listando produtos
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("->", product.Name, "R$", product.Price)
		}
	}
}
