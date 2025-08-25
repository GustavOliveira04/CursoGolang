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
	gorm.Model
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Product{})

	// Inserindo dados
	products := []Product{
		{Name: "Notebook", Price: 5000.00},
		{Name: "Ps5", Price: 3500.00},
	}
	db.Create(&products)

	// Atualizando produto
	var p Product
	db.First(&p, 1)
	p.Name = "Xbox Series X2"
	db.Save(&p)

	// Excluindo produto
	var p2 Product
	db.First(&p2, 5)
	db.Delete(&p2)
}
