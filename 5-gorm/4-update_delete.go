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

	// Update de um produto
	// var p Product
	// db.First(&p, 1)
	// p.Name = "Data Show"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)

	// Atualizando múltiplas colunas
	db.Model(&Product{}).Where("id = ?", 1).Updates(Product{Price: 5000, Name: "Xbox Series X"})

	// Delete de produtos
	var p3 Product
	db.First(&p3, 4)
	db.Delete(&p3)
}
