package main

import (
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

	// Migrar a tabela
	db.AutoMigrate(&Product{})

	// Inserção de dados de um a um
	db.Create(&Product{Name: "PS5", Price: 3500})

	// Inserção de dados com um slice (quantos quiser)
	products := []Product{
		{Name: "Notebook", Price: 5000},
		{Name: "Tablet", Price: 1500},
		{Name: "Fogão", Price: 1000},
	}
	db.Create(&products)

	
}
