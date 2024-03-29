package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product defines a product
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product

	db.First(&product, 1)
	log.Printf("%+v\n", product)

	db.First(&product, "code = ?", "L1212")
	log.Printf("%+v\n", product)

	db.Model(&product).Update("Price", 2000)
	db.First(&product, "code = ?", "L1212")
	log.Printf("%+v\n", product)

	db.Delete(&product)
}
