//Initial migration to version: 0.1.0

package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Item struct {
	gorm.Model         //`json:"model"`
	Title      string  //`json:"name"`
	Price      float32 //`json:"price"`
}

func main() {

	// Connect to DB
	dsn := "host=oleg-web.devops.rebrain.srwx.net user=api password=GQt5MTyVPuf9vsVWoWDT9YCn dbname=api_test port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate to the schema
	// version: 0.1.0

	// Add Items Table
	s, _ := schema.Parse(&Item{}, &sync.Map{}, schema.NamingStrategy{})
	if db.Migrator().HasTable(&Item{}) {
		fmt.Println("Info: Table '" + s.Table + "' already exists.\nMigration skipped!")
	} else {
		db.AutoMigrate(&Item{})
		fmt.Println("Table '" + s.Table + "' created.")
		// Add sample data
		var items = []Item{{Title: "apple", Price: 100}, {Title: "orange", Price: 120}, {Title: "banana", Price: 95.5}}
		db.Create(&items)
		for _, item := range items {
			fmt.Println(item.ID, item.Title, item.Price)
		}
		fmt.Println("Migration successfully finished!")
	}
}
