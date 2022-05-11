//Migration to version: 0.2.0

package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Warehouse struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
}

type Item struct {
	gorm.Model  `json:"model"`
	Title       string    `json:"name"`
	Price       float32   `json:"price"`
	WarehouseID int       `gorm:"default:1"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func main() {

	// Connect to DB
	dsn := "host=oleg-web.devops.rebrain.srwx.net user=api password=GQt5MTyVPuf9vsVWoWDT9YCn dbname=api_test port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// version: 0.2.0

	// Add Warehouse Table
	s, _ := schema.Parse(&Warehouse{}, &sync.Map{}, schema.NamingStrategy{})
	if db.Migrator().HasTable(&Warehouse{}) {
		fmt.Println("...Info: Table '" + s.Table + "' already exists.\nMigration skipped!")
	} else {
		db.AutoMigrate(&Warehouse{})
		fmt.Println("Table '" + s.Table + "' created.")
		// Add sample data to Warehouse Table
		var warehouses = []Warehouse{{Name: "Base"}, {Name: "West"}, {Name: "East"}, {Name: "North"}}
		db.Create(&warehouses)
		for _, warehouse := range warehouses {
			fmt.Println(warehouse.ID, warehouse.Name)
		}
	}

	// Check Table Item
	if !(db.Migrator().HasTable(&Item{})) {
		db.AutoMigrate(&Item{})
		fmt.Println("Table '" + s.Table + "' created.")
		// Add sample data
		var items = []Item{{Title: "apple", Price: 100}, {Title: "orange", Price: 120}, {Title: "banana", Price: 95.5}}
		db.Create(&items)
		for _, item := range items {
			fmt.Println(item.ID, item.Title, item.Price)
		}
	}
	// Add WarehouseID Column to Item Table
	if db.Migrator().HasColumn(&Item{}, "WarehouseID") {
		fmt.Println("...Info: Table '" + s.Table + "' already has Column 'WarehouseID'.\nMigration skipped!")
	} else {
		db.Migrator().AddColumn(&Item{}, "WarehouseID")
		fmt.Println("Column 'WarehouseID' added to Table '" + s.Table + "'.")
	}
	fmt.Println("Migration successfully finished!")
}
