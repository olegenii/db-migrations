//Initial migration to version: 0.1.0

package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "fmt"
)

type Item struct {
	gorm.Model		//`json:"model"`
	Title string 	//`json:"name"`
	Price float32 	//`json:"price"`
}

func main() {
	
	// Connect to DB
	dsn := "host=oleg-web.devops.rebrain.srwx.net user=api password=GQt5MTyVPuf9vsVWoWDT9YCn dbname=api port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	// version: 0.1.0

	db.AutoMigrate(&Item{})

	// Add sample data
	var items = []Item{{Title: "apple", Price: 100}, {Title: "orange", Price: 120}, {Title: "banana", Price: 95.5}}
	db.Create(&items)
	for _, item := range items {
		fmt.Println(item.ID, item.Title, item.Price)
	}
}