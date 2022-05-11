//Migration to version: 0.2.0

package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "fmt"
)

type Warehouse struct {
	gorm.Model					//`json:"model"`
	Name 			string 		//`json:"name"`
}

type Item struct {
	gorm.Model					//`json:"model"`
	Title 			string 		//`json:"name"`
	Price 			float32 	//`json:"price"`
	WarehouseID 	int
	Warehouse		Warehouse	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func main() {
	
	// Connect to DB
	dsn := "host=oleg-web.devops.rebrain.srwx.net user=api password=GQt5MTyVPuf9vsVWoWDT9YCn dbname=api port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  

	// Migrate the schema
	// version: 0.2.0
	
	// Add Warehouse Table
	db.AutoMigrate(&Warehouse{})

	// Add sample data to Warehouse Table
	var warehouses = []Item{{Name: "West"}, {Name: "East"}, {name: "North"}}
	db.Create(&items)
	for _, item := range items {
		fmt.Println(item.ID, item.Name)
	}

	//Add WarehouseID Column to Item Table
	db.Migrator().AddColumn(&Item{}, "WarehouseID")
}